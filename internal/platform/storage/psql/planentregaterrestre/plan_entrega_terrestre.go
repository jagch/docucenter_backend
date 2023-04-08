package terrestre

import (
	"context"
	"database/sql"
	"fmt"
	gb "jagch/backend/internal/global"
	planentrega "jagch/backend/internal/planentrega"
	"strconv"
	"strings"
)

type StorageTerrestre struct {
	db *sql.DB
}

func NewPETerrestreStorage(db *sql.DB) *StorageTerrestre {
	return &StorageTerrestre{
		db: db,
	}
}

func (r *StorageTerrestre) Create(ctx context.Context, pet any) (string, error) {
	stmt, err := r.db.PrepareContext(ctx, `
		INSERT INTO public.plan_entrega_terrestre(
			id, 
			tipo_producto, 
			cantidad_producto, 
			fecha_registro, 
			fecha_entrega, 
			id_bodega_entrega, 
			precio_envio, 
			placa_vehiculo, 
			nro_guia,
			id_cliente,
			dscto
		)VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id`)
	if err != nil {
		return "", fmt.Errorf("error al insertar el plan de entrega terrestre en la base de datos: %v", err)
	}

	defer stmt.Close()

	peTerrestre := pet.(planentrega.PETerrestre)
	var id string
	err = stmt.QueryRowContext(ctx,
		peTerrestre.ID.String(),
		peTerrestre.PlanEntrega.TipoProducto.String(),
		peTerrestre.PlanEntrega.CantidadProducto.Int(),
		gb.CurrentTime,
		peTerrestre.PlanEntrega.FechaEntrega.Date(),
		peTerrestre.IDBodegaEntrega.Int(),
		peTerrestre.PlanEntrega.PrecioEnvio.Float64(),
		peTerrestre.PlacaVehiculo.String(),
		peTerrestre.PlanEntrega.NroGuia.String(),
		peTerrestre.PlanEntrega.IDCliente.String(),
		peTerrestre.PlanEntrega.Dscto.Float64(),
	).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("error al insertar el plan de entrega terrestre en la base de datos: %v", err)
	}

	return id, nil
}

func (r *StorageTerrestre) GetAll(ctx context.Context) (any, error) {
	stmt, err := r.db.PrepareContext(
		ctx,
		`	SELECT 
				pet.id, 
				pet.tipo_producto, 
				pet.cantidad_producto, 
				pet.fecha_registro, 
				pet.fecha_entrega, 
				pet.id_bodega_entrega,
				b.nombre, 
				pet.precio_envio, 
				pet.placa_vehiculo, 
				pet.nro_guia, 
				pet.id_cliente, 
				pet.dscto
			FROM plan_entrega_terrestre pet
			INNER JOIN bodega b on pet.id_bodega_entrega = b.id
		`)
	if err != nil {
		return nil, fmt.Errorf("error intentando obtener los planes de entrega terrestre de la base de datos: %v", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("error intentando obtener los planes de entrega terrestre de la base de datos: %v", err)
	}
	defer rows.Close()

	var planesEntregaTerrestresReqRes planentrega.PlanesEntregaTerrestresReqRes
	i := 1
	for rows.Next() {
		var planEntregaTerrestreReqRes planentrega.PlanEntregaTerrestreReqRes
		if err = rows.Scan(
			&planEntregaTerrestreReqRes.ID,
			&planEntregaTerrestreReqRes.TipoProducto,
			&planEntregaTerrestreReqRes.CantidadProducto,
			&planEntregaTerrestreReqRes.FechaRegistro,
			&planEntregaTerrestreReqRes.FechaEntrega,
			&planEntregaTerrestreReqRes.IDBodegaEntrega,
			&planEntregaTerrestreReqRes.BodegaEntrega,
			&planEntregaTerrestreReqRes.PrecioEnvio,
			&planEntregaTerrestreReqRes.PlacaVehiculo,
			&planEntregaTerrestreReqRes.NroGuia,
			&planEntregaTerrestreReqRes.IDCliente,
			&planEntregaTerrestreReqRes.Dscto,
		); err != nil {
			return nil, fmt.Errorf("error intentando obtener los planes de entrega terrestre de la base de datos: %v", err)
		}
		planEntregaTerrestreReqRes.Key = i
		planesEntregaTerrestresReqRes = append(planesEntregaTerrestresReqRes, planEntregaTerrestreReqRes)

		i++
	}

	return planesEntregaTerrestresReqRes, nil
}

func (r *StorageTerrestre) Update(ctx context.Context, pet any) (string, error) {
	stmt, err := r.db.PrepareContext(ctx, `
		UPDATE public.plan_entrega_terrestre SET 
			tipo_producto = $1, 
			cantidad_producto = $2, 
			fecha_registro = $3, 
			fecha_entrega = $4, 
			id_bodega_entrega = $5, 
			precio_envio = $6, 
			placa_vehiculo = $7, 
			nro_guia = $8,
			id_cliente = $9,
			dscto = $10
		WHERE 
			id = $11
		RETURNING id`)
	if err != nil {
		return "", fmt.Errorf("error al insertar el plan de entrega terrestre en la base de datos: %v", err)
	}
	defer stmt.Close()

	peTerrestre := pet.(planentrega.PETerrestre)

	var id string
	err = stmt.QueryRowContext(ctx,
		peTerrestre.PlanEntrega.TipoProducto.String(),
		peTerrestre.PlanEntrega.CantidadProducto.Int(),
		gb.CurrentTime,
		peTerrestre.PlanEntrega.FechaEntrega.Date(),
		peTerrestre.IDBodegaEntrega.Int(),
		peTerrestre.PlanEntrega.PrecioEnvio.Float64(),
		peTerrestre.PlacaVehiculo.String(),
		peTerrestre.PlanEntrega.NroGuia.String(),
		peTerrestre.PlanEntrega.IDCliente.String(),
		peTerrestre.PlanEntrega.Dscto.Float64(),
		peTerrestre.ID.String(),
	).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("error al actualizar el plan de entrega terrestre en la base de datos: %v", err)
	}

	return id, nil
}

func (r *StorageTerrestre) Delete(ctx context.Context, id string) error {
	stmt, err := r.db.PrepareContext(ctx, `DELETE FROM public.plan_entrega_terrestre WHERE id = $1 RETURNING id`)
	if err != nil {
		return fmt.Errorf("error al eliminar el plan de entrega terrestre en la base de datos: %v", err)
	}
	defer stmt.Close()

	idPET := ""
	err = stmt.QueryRowContext(ctx, id).Scan(&idPET)
	if err != nil {
		return fmt.Errorf("error al eliminar el plan de entrega terrestre en la base de datos: %v", err)
	}

	if idPET == "" {
		return fmt.Errorf("error al eliminar el plan de entrega terrestre en la base de datos")
	}

	return nil
}

func (r *StorageTerrestre) Search(ctx context.Context, mapSearchParams map[string]planentrega.SearchParams, page int) (any, error) {
	query, args, err := buildSQLQuery(mapSearchParams, page)
	if err != nil {
		return nil, fmt.Errorf("error con la consulta a la base de datos %v", err)
	}

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error al realizar la búsqueda en la base de datos: %v", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, fmt.Errorf("error intentando obtener los planes de entrega terrestre de la base de datos: %v", err)
	}
	defer rows.Close()

	var planesEntregaTerrestresReqRes planentrega.PlanesEntregaTerrestresReqRes
	i := 1
	for rows.Next() {
		var planEntregaTerrestreReqRes planentrega.PlanEntregaTerrestreReqRes
		if err = rows.Scan(
			&planEntregaTerrestreReqRes.ID,
			&planEntregaTerrestreReqRes.TipoProducto,
			&planEntregaTerrestreReqRes.CantidadProducto,
			&planEntregaTerrestreReqRes.FechaRegistro,
			&planEntregaTerrestreReqRes.FechaEntrega,
			&planEntregaTerrestreReqRes.IDBodegaEntrega,
			&planEntregaTerrestreReqRes.BodegaEntrega,
			&planEntregaTerrestreReqRes.PrecioEnvio,
			&planEntregaTerrestreReqRes.PlacaVehiculo,
			&planEntregaTerrestreReqRes.NroGuia,
			&planEntregaTerrestreReqRes.IDCliente,
			&planEntregaTerrestreReqRes.Dscto,
		); err != nil {
			return nil, fmt.Errorf("error al intentar obtener los clientes de la base de datos: %v", err)
		}
		planEntregaTerrestreReqRes.Key = i
		planesEntregaTerrestresReqRes = append(planesEntregaTerrestresReqRes, planEntregaTerrestreReqRes)

		i++
	}

	return planesEntregaTerrestresReqRes, nil
}

func buildSQLQuery(mapSearchParams map[string]planentrega.SearchParams, page int) (string, []any, error) {
	offset := (page - 1) * gb.PerPage

	query := `
			SELECT 
				pet.id, 
				pet.tipo_producto, 
				pet.cantidad_producto, 
				pet.fecha_registro, 
				pet.fecha_entrega, 
				pet.id_bodega_entrega,
				b.nombre, 
				pet.precio_envio, 
				pet.placa_vehiculo, 
				pet.nro_guia, 
				pet.id_cliente, 
				pet.dscto
			FROM plan_entrega_terrestre pet
			INNER JOIN bodega b on pet.id_bodega_entrega = b.id
			WHERE `

	counter := 0
	if mapSearchParams[gb.FieldIDCliente].Active {
		query = fmt.Sprintf("%s pet.%s = $X AND (pet.%s IS NULL OR pet.%s IS NOT NULL) AND ", query, mapSearchParams[gb.FieldIDCliente].NameParam, mapSearchParams[gb.FieldIDCliente].NameParam, mapSearchParams[gb.FieldIDCliente].NameParam)

		counter++
	}

	if mapSearchParams[gb.FieldFechaEntrega].Active {
		query = fmt.Sprintf("%s pet.%s = $X AND (pet.%s IS NULL OR pet.%s IS NOT NULL) AND ", query, mapSearchParams[gb.FieldFechaEntrega].NameParam, mapSearchParams[gb.FieldFechaEntrega].NameParam, mapSearchParams[gb.FieldFechaEntrega].NameParam)

		counter++
	}

	if mapSearchParams[gb.FieldPlacaVehiculo].Active {
		query = fmt.Sprintf("%s pet.%s = $X AND (pet.%s IS NULL OR pet.%s IS NOT NULL) ", query, mapSearchParams[gb.FieldPlacaVehiculo].NameParam, mapSearchParams[gb.FieldPlacaVehiculo].NameParam, mapSearchParams[gb.FieldPlacaVehiculo].NameParam)

		counter++
	}

	if counter == 0 {
		return "", nil, fmt.Errorf("%s", "Al menos debe haber un parámetro para la búsqueda")
	}

	query = enumerateDolarSign(query)

	query = fmt.Sprintf("%s LIMIT %d OFFSET %d", query, gb.PerPage, offset)

	query = deleteFinalANDBeforeLIMIT(query)

	var args []any
	for _, v := range mapSearchParams {
		if v.Active {
			args = append(args, v.Value)
		}
	}

	return query, args, nil
}

func enumerateDolarSign(query string) string {
	arrQuery := strings.Split(query, "X")
	lenArrStr := len(arrQuery)

	i := 0
	finalQuery := ""
	for i < lenArrStr {
		if strings.Contains(arrQuery[i], "$") {
			finalQuery += string(arrQuery[i]) + strconv.Itoa(i+1)
		} else {
			finalQuery += string(arrQuery[i])
		}

		i++
	}

	return finalQuery
}

func deleteFinalANDBeforeLIMIT(query string) string {
	if strings.Contains(query, "AND  LIMIT") {
		return strings.Replace(query, "AND  LIMIT", "LIMIT", 1)
	}

	return query
}
