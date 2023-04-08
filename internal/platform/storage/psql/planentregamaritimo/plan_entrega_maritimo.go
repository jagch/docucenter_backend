package maritimo

import (
	"context"
	"database/sql"
	"fmt"
	gb "jagch/backend/internal/global"
	planentrega "jagch/backend/internal/planentrega"
	"strconv"
	"strings"
)

type StorageMaritimo struct {
	db *sql.DB
}

func NewPEMaritimoStorage(db *sql.DB) *StorageMaritimo {
	return &StorageMaritimo{
		db: db,
	}
}

func (r *StorageMaritimo) Create(ctx context.Context, pem any) (string, error) {
	stmt, err := r.db.PrepareContext(ctx, `
		INSERT INTO public.plan_entrega_maritimo(
			id, 
			tipo_producto, 
			cantidad_producto, 
			fecha_registro, 
			fecha_entrega, 
			id_puerto_entrega, 
			precio_envio, 
			nro_flota, 
			nro_guia,
			id_cliente,
			dscto
		)VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id`)
	if err != nil {
		return "", fmt.Errorf("error al insertar el plan de entrega maritimo en la base de datos: %v", err)
	}

	defer stmt.Close()

	peMaritimo := pem.(planentrega.PEMaritimo)
	var id string
	err = stmt.QueryRowContext(ctx,
		peMaritimo.ID.String(),
		peMaritimo.PlanEntrega.TipoProducto.String(),
		peMaritimo.PlanEntrega.CantidadProducto.Int(),
		gb.CurrentTime,
		peMaritimo.PlanEntrega.FechaEntrega.Date(),
		peMaritimo.IDPuertoEntrega.Int(),
		peMaritimo.PlanEntrega.PrecioEnvio.Float64(),
		peMaritimo.NroFlota.String(),
		peMaritimo.PlanEntrega.NroGuia.String(),
		peMaritimo.PlanEntrega.IDCliente.String(),
		peMaritimo.PlanEntrega.Dscto.Float64(),
	).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("error al insertar el plan de entrega maritimo en la base de datos: %v", err)
	}

	return id, nil
}

func (r *StorageMaritimo) GetAll(ctx context.Context) (any, error) {
	stmt, err := r.db.PrepareContext(
		ctx,
		`	SELECT 
				pem.id, 
				pem.tipo_producto, 
				pem.cantidad_producto, 
				pem.fecha_registro, 
				pem.fecha_entrega, 
				pem.id_puerto_entrega,
				b.nombre, 
				pem.precio_envio, 
				pem.nro_flota, 
				pem.nro_guia, 
				pem.id_cliente, 
				pem.dscto
			FROM plan_entrega_maritimo pem
			INNER JOIN puerto b on pem.id_puerto_entrega = b.id
		`)
	if err != nil {
		return nil, fmt.Errorf("error intentando obtener los planes de entrega maritimo de la base de datos: %v", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("error intentando obtener los planes de entrega maritimo de la base de datos: %v", err)
	}
	defer rows.Close()

	var planesEntregaMaritimosReqRes planentrega.PlanesEntregaMaritimosReqRes
	i := 1
	for rows.Next() {
		var planEntregaMaritimoReqRes planentrega.PlanEntregaMaritimoReqRes
		if err = rows.Scan(
			&planEntregaMaritimoReqRes.ID,
			&planEntregaMaritimoReqRes.TipoProducto,
			&planEntregaMaritimoReqRes.CantidadProducto,
			&planEntregaMaritimoReqRes.FechaRegistro,
			&planEntregaMaritimoReqRes.FechaEntrega,
			&planEntregaMaritimoReqRes.IDPuertoEntrega,
			&planEntregaMaritimoReqRes.PuertoEntrega,
			&planEntregaMaritimoReqRes.PrecioEnvio,
			&planEntregaMaritimoReqRes.NroFlota,
			&planEntregaMaritimoReqRes.NroGuia,
			&planEntregaMaritimoReqRes.IDCliente,
			&planEntregaMaritimoReqRes.Dscto,
		); err != nil {
			return nil, fmt.Errorf("error intentando obtener los planes de entrega maritimo de la base de datos: %v", err)
		}
		planEntregaMaritimoReqRes.Key = i
		planesEntregaMaritimosReqRes = append(planesEntregaMaritimosReqRes, planEntregaMaritimoReqRes)

		i++
	}

	return planesEntregaMaritimosReqRes, nil
}

func (r *StorageMaritimo) Update(ctx context.Context, pem any) (string, error) {
	stmt, err := r.db.PrepareContext(ctx, `
		UPDATE public.plan_entrega_maritimo SET 
			tipo_producto = $1, 
			cantidad_producto = $2, 
			fecha_registro = $3, 
			fecha_entrega = $4, 
			id_puerto_entrega = $5, 
			precio_envio = $6, 
			nro_flota = $7, 
			nro_guia = $8,
			id_cliente = $9,
			dscto = $10
		WHERE 
			id = $11
		RETURNING id`)
	if err != nil {
		return "", fmt.Errorf("error al insertar el plan de entrega maritimo en la base de datos: %v", err)
	}
	defer stmt.Close()

	peMaritimo := pem.(planentrega.PEMaritimo)

	var id string
	err = stmt.QueryRowContext(ctx,
		peMaritimo.PlanEntrega.TipoProducto.String(),
		peMaritimo.PlanEntrega.CantidadProducto.Int(),
		gb.CurrentTime,
		peMaritimo.PlanEntrega.FechaEntrega.Date(),
		peMaritimo.IDPuertoEntrega.Int(),
		peMaritimo.PlanEntrega.PrecioEnvio.Float64(),
		peMaritimo.NroFlota.String(),
		peMaritimo.PlanEntrega.NroGuia.String(),
		peMaritimo.PlanEntrega.IDCliente.String(),
		peMaritimo.PlanEntrega.Dscto.Float64(),
		peMaritimo.ID.String(),
	).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("error al actualizar el plan de entrega maritimo en la base de datos: %v", err)
	}

	return id, nil
}

func (r *StorageMaritimo) Delete(ctx context.Context, id string) error {
	stmt, err := r.db.PrepareContext(ctx, `DELETE FROM public.plan_entrega_maritimo WHERE id = $1 RETURNING id`)
	if err != nil {
		return fmt.Errorf("error al eliminar el plan de entrega maritimo en la base de datos: %v", err)
	}
	defer stmt.Close()

	idPEM := ""
	err = stmt.QueryRowContext(ctx, id).Scan(&idPEM)
	if err != nil {
		return fmt.Errorf("error al eliminar el plan de entrega maritimo en la base de datos: %v", err)
	}

	if idPEM == "" {
		return fmt.Errorf("error al eliminar el plan de entrega maritimo en la base de datos")
	}

	return nil
}

func (r *StorageMaritimo) Search(ctx context.Context, mapSearchParams map[string]planentrega.SearchParams, page int) (any, error) {
	query, args, err := buildSQLQuery(mapSearchParams, page)
	if err != nil {
		return nil, fmt.Errorf("error con la consulta a la base de datos %v", err)
	}

	fmt.Println(query)

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error al realizar la búsqueda en la base de datos: %v", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, fmt.Errorf("error intentando obtener los planes de entrega maritimo de la base de datos: %v", err)
	}
	defer rows.Close()

	var planesEntregaMaritimosReqRes planentrega.PlanesEntregaMaritimosReqRes
	i := 1
	for rows.Next() {
		var planEntregaMaritimoReqRes planentrega.PlanEntregaMaritimoReqRes
		if err = rows.Scan(
			&planEntregaMaritimoReqRes.ID,
			&planEntregaMaritimoReqRes.TipoProducto,
			&planEntregaMaritimoReqRes.CantidadProducto,
			&planEntregaMaritimoReqRes.FechaRegistro,
			&planEntregaMaritimoReqRes.FechaEntrega,
			&planEntregaMaritimoReqRes.IDPuertoEntrega,
			&planEntregaMaritimoReqRes.PuertoEntrega,
			&planEntregaMaritimoReqRes.PrecioEnvio,
			&planEntregaMaritimoReqRes.NroFlota,
			&planEntregaMaritimoReqRes.NroGuia,
			&planEntregaMaritimoReqRes.IDCliente,
			&planEntregaMaritimoReqRes.Dscto,
		); err != nil {
			return nil, fmt.Errorf("error al intentar obtener los clientes de la base de datos: %v", err)
		}
		planEntregaMaritimoReqRes.Key = i
		planesEntregaMaritimosReqRes = append(planesEntregaMaritimosReqRes, planEntregaMaritimoReqRes)

		i++
	}

	return planesEntregaMaritimosReqRes, nil
}

func buildSQLQuery(mapSearchParams map[string]planentrega.SearchParams, page int) (string, []any, error) {
	offset := (page - 1) * gb.PerPage

	query := `
			SELECT 
				pem.id, 
				pem.tipo_producto, 
				pem.cantidad_producto, 
				pem.fecha_registro, 
				pem.fecha_entrega, 
				pem.id_puerto_entrega,
				b.nombre, 
				pem.precio_envio, 
				pem.nro_flota, 
				pem.nro_guia, 
				pem.id_cliente, 
				pem.dscto
			FROM plan_entrega_maritimo pem
			INNER JOIN puerto b on pem.id_puerto_entrega = b.id
			WHERE `

	counter := 0
	if mapSearchParams[gb.FieldIDCliente].Active {
		query = fmt.Sprintf("%s pem.%s = $X AND (pem.%s IS NULL OR pem.%s IS NOT NULL) AND ", query, mapSearchParams[gb.FieldIDCliente].NameParam, mapSearchParams[gb.FieldIDCliente].NameParam, mapSearchParams[gb.FieldIDCliente].NameParam)

		counter++
	}

	if mapSearchParams[gb.FieldFechaEntrega].Active {
		query = fmt.Sprintf("%s pem.%s = $X AND (pem.%s IS NULL OR pem.%s IS NOT NULL) AND ", query, mapSearchParams[gb.FieldFechaEntrega].NameParam, mapSearchParams[gb.FieldFechaEntrega].NameParam, mapSearchParams[gb.FieldFechaEntrega].NameParam)

		counter++
	}

	if mapSearchParams[gb.FieldNroFlota].Active {
		query = fmt.Sprintf("%s pem.%s = $X AND (pem.%s IS NULL OR pem.%s IS NOT NULL) ", query, mapSearchParams[gb.FieldNroFlota].NameParam, mapSearchParams[gb.FieldNroFlota].NameParam, mapSearchParams[gb.FieldNroFlota].NameParam)

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
