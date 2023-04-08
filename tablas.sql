--
-- PostgreSQL database dump
--

-- Dumped from database version 12.3
-- Dumped by pg_dump version 12.3

-- Started on 2023-04-07 18:06:09

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 3 (class 2615 OID 2200)
-- Name: public; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA IF NOT EXISTS public;


ALTER SCHEMA public OWNER TO postgres;

--
-- TOC entry 2876 (class 0 OID 0)
-- Dependencies: 3
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 208 (class 1259 OID 1913146)
-- Name: bodega; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.bodega (
    id integer NOT NULL,
    nombre text
);


ALTER TABLE public.bodega OWNER TO postgres;

--
-- TOC entry 207 (class 1259 OID 1913144)
-- Name: bodegas_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.bodegas_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.bodegas_id_seq OWNER TO postgres;

--
-- TOC entry 2877 (class 0 OID 0)
-- Dependencies: 207
-- Name: bodegas_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.bodegas_id_seq OWNED BY public.bodega.id;


--
-- TOC entry 202 (class 1259 OID 1896745)
-- Name: cliente; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.cliente (
    id text NOT NULL,
    nombre text NOT NULL
);


ALTER TABLE public.cliente OWNER TO postgres;

--
-- TOC entry 206 (class 1259 OID 1913136)
-- Name: plan_entrega_maritimo; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.plan_entrega_maritimo (
    id text NOT NULL,
    tipo_producto text,
    cantidad_producto text,
    fecha_registro timestamp without time zone,
    fecha_entrega timestamp without time zone,
    id_puerto_entrega integer,
    precio_envio numeric(11,2),
    nro_flota text,
    nro_guia text,
    id_cliente text,
    dscto numeric(11,2)
);


ALTER TABLE public.plan_entrega_maritimo OWNER TO postgres;

--
-- TOC entry 205 (class 1259 OID 1913128)
-- Name: plan_entrega_terrestre; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.plan_entrega_terrestre (
    id text NOT NULL,
    tipo_producto text,
    cantidad_producto text,
    fecha_registro timestamp without time zone,
    fecha_entrega timestamp without time zone,
    id_bodega_entrega integer,
    precio_envio numeric(11,2),
    placa_vehiculo text,
    nro_guia text,
    id_cliente text,
    dscto numeric(11,2)
);


ALTER TABLE public.plan_entrega_terrestre OWNER TO postgres;

--
-- TOC entry 210 (class 1259 OID 1913157)
-- Name: puerto; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.puerto (
    id integer NOT NULL,
    nombre text
);


ALTER TABLE public.puerto OWNER TO postgres;

--
-- TOC entry 209 (class 1259 OID 1913155)
-- Name: puerto_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.puerto_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.puerto_id_seq OWNER TO postgres;

--
-- TOC entry 2878 (class 0 OID 0)
-- Dependencies: 209
-- Name: puerto_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.puerto_id_seq OWNED BY public.puerto.id;


--
-- TOC entry 203 (class 1259 OID 1896753)
-- Name: usuario; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.usuario (
    usuario text NOT NULL,
    clave text NOT NULL,
    id smallint NOT NULL
);


ALTER TABLE public.usuario OWNER TO postgres;

--
-- TOC entry 204 (class 1259 OID 1896761)
-- Name: usuario_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.usuario_id_seq
    AS smallint
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.usuario_id_seq OWNER TO postgres;

--
-- TOC entry 2879 (class 0 OID 0)
-- Dependencies: 204
-- Name: usuario_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.usuario_id_seq OWNED BY public.usuario.id;


--
-- TOC entry 2718 (class 2604 OID 1913149)
-- Name: bodega id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bodega ALTER COLUMN id SET DEFAULT nextval('public.bodegas_id_seq'::regclass);


--
-- TOC entry 2719 (class 2604 OID 1913160)
-- Name: puerto id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.puerto ALTER COLUMN id SET DEFAULT nextval('public.puerto_id_seq'::regclass);


--
-- TOC entry 2717 (class 2604 OID 1896763)
-- Name: usuario id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.usuario ALTER COLUMN id SET DEFAULT nextval('public.usuario_id_seq'::regclass);


--
-- TOC entry 2868 (class 0 OID 1913146)
-- Dependencies: 208
-- Data for Name: bodega; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.bodega VALUES (1, 'Bodega A');
INSERT INTO public.bodega VALUES (2, 'Bodega B');
INSERT INTO public.bodega VALUES (3, 'Bodega C');


--
-- TOC entry 2862 (class 0 OID 1896745)
-- Dependencies: 202
-- Data for Name: cliente; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 2866 (class 0 OID 1913136)
-- Dependencies: 206
-- Data for Name: plan_entrega_maritimo; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 2865 (class 0 OID 1913128)
-- Dependencies: 205
-- Data for Name: plan_entrega_terrestre; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 2870 (class 0 OID 1913157)
-- Dependencies: 210
-- Data for Name: puerto; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.puerto VALUES (1, 'Puerto A');
INSERT INTO public.puerto VALUES (2, 'Puerto B');
INSERT INTO public.puerto VALUES (3, 'Puerto C');


--
-- TOC entry 2863 (class 0 OID 1896753)
-- Dependencies: 203
-- Data for Name: usuario; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.usuario VALUES ('admin', '123', 1);
INSERT INTO public.usuario VALUES ('1', 'admin', 123);


--
-- TOC entry 2880 (class 0 OID 0)
-- Dependencies: 207
-- Name: bodegas_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.bodegas_id_seq', 3, true);


--
-- TOC entry 2881 (class 0 OID 0)
-- Dependencies: 209
-- Name: puerto_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.puerto_id_seq', 3, true);


--
-- TOC entry 2882 (class 0 OID 0)
-- Dependencies: 204
-- Name: usuario_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.usuario_id_seq', 1, true);


--
-- TOC entry 2729 (class 2606 OID 1913154)
-- Name: bodega bodegas_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bodega
    ADD CONSTRAINT bodegas_pkey PRIMARY KEY (id);


--
-- TOC entry 2721 (class 2606 OID 1896752)
-- Name: cliente cliente_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cliente
    ADD CONSTRAINT cliente_pkey PRIMARY KEY (id);


--
-- TOC entry 2727 (class 2606 OID 1913143)
-- Name: plan_entrega_maritimo plan_entrega_maritimo_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.plan_entrega_maritimo
    ADD CONSTRAINT plan_entrega_maritimo_pkey PRIMARY KEY (id);


--
-- TOC entry 2725 (class 2606 OID 1913135)
-- Name: plan_entrega_terrestre plan_entrega_terrestre_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.plan_entrega_terrestre
    ADD CONSTRAINT plan_entrega_terrestre_pkey PRIMARY KEY (id);


--
-- TOC entry 2731 (class 2606 OID 1913165)
-- Name: puerto puerto_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.puerto
    ADD CONSTRAINT puerto_pkey PRIMARY KEY (id);


--
-- TOC entry 2723 (class 2606 OID 1896771)
-- Name: usuario usuario_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.usuario
    ADD CONSTRAINT usuario_pkey PRIMARY KEY (id);


--
-- TOC entry 2732 (class 2606 OID 1913182)
-- Name: plan_entrega_terrestre fk_id_bodega_entrega; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.plan_entrega_terrestre
    ADD CONSTRAINT fk_id_bodega_entrega FOREIGN KEY (id_bodega_entrega) REFERENCES public.bodega(id) NOT VALID;


--
-- TOC entry 2733 (class 2606 OID 1913187)
-- Name: plan_entrega_terrestre fk_id_cliente; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.plan_entrega_terrestre
    ADD CONSTRAINT fk_id_cliente FOREIGN KEY (id_cliente) REFERENCES public.cliente(id) NOT VALID;


--
-- TOC entry 2735 (class 2606 OID 1913197)
-- Name: plan_entrega_maritimo fk_id_cliente; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.plan_entrega_maritimo
    ADD CONSTRAINT fk_id_cliente FOREIGN KEY (id_cliente) REFERENCES public.cliente(id) NOT VALID;


--
-- TOC entry 2734 (class 2606 OID 1913192)
-- Name: plan_entrega_maritimo fk_id_puerto_entrega; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.plan_entrega_maritimo
    ADD CONSTRAINT fk_id_puerto_entrega FOREIGN KEY (id_puerto_entrega) REFERENCES public.puerto(id) NOT VALID;


-- Completed on 2023-04-07 18:06:09

--
-- PostgreSQL database dump complete
--

