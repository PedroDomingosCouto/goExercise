--
-- PostgreSQL database dump
--

-- Dumped from database version 13.1
-- Dumped by pg_dump version 13.4

-- Started on 2021-09-20 19:39:24 WEST

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
-- TOC entry 200 (class 1259 OID 16710)
-- Name: house_id_seq; Type: SEQUENCE; Schema: public; Owner: pedro
--

CREATE SEQUENCE public.house_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    MAXVALUE 2147483647
    CACHE 1;


ALTER TABLE public.house_id_seq OWNER TO pedro;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 201 (class 1259 OID 16712)
-- Name: house; Type: TABLE; Schema: public; Owner: pedro
--

CREATE TABLE public.house (
    id integer DEFAULT nextval('public.house_id_seq'::regclass) NOT NULL,
    name character varying NOT NULL,
    animal character varying NOT NULL,
    motto character varying(250) NOT NULL
);


ALTER TABLE public.house OWNER TO pedro;

--
-- TOC entry 202 (class 1259 OID 16721)
-- Name: person_id_seq; Type: SEQUENCE; Schema: public; Owner: pedro
--

CREATE SEQUENCE public.person_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    MAXVALUE 2147483647
    CACHE 1;


ALTER TABLE public.person_id_seq OWNER TO pedro;

--
-- TOC entry 203 (class 1259 OID 16738)
-- Name: person; Type: TABLE; Schema: public; Owner: pedro
--

CREATE TABLE public.person (
    id integer DEFAULT nextval('public.person_id_seq'::regclass) NOT NULL,
    houseid integer NOT NULL,
    name character varying NOT NULL,
    ismarried boolean DEFAULT false
);


ALTER TABLE public.person OWNER TO pedro;

--
-- TOC entry 3262 (class 0 OID 16712)
-- Dependencies: 201
-- Data for Name: house; Type: TABLE DATA; Schema: public; Owner: pedro
--

COPY public.house (id, name, animal, motto) FROM stdin;
\.


--
-- TOC entry 3264 (class 0 OID 16738)
-- Dependencies: 203
-- Data for Name: person; Type: TABLE DATA; Schema: public; Owner: pedro
--

COPY public.person (id, houseid, name, ismarried) FROM stdin;
\.


--
-- TOC entry 3270 (class 0 OID 0)
-- Dependencies: 200
-- Name: house_id_seq; Type: SEQUENCE SET; Schema: public; Owner: pedro
--

SELECT pg_catalog.setval('public.house_id_seq', 27, true);


--
-- TOC entry 3271 (class 0 OID 0)
-- Dependencies: 202
-- Name: person_id_seq; Type: SEQUENCE SET; Schema: public; Owner: pedro
--

SELECT pg_catalog.setval('public.person_id_seq', 14, true);


--
-- TOC entry 3127 (class 2606 OID 16720)
-- Name: house house_pk; Type: CONSTRAINT; Schema: public; Owner: pedro
--

ALTER TABLE ONLY public.house
    ADD CONSTRAINT house_pk PRIMARY KEY (id);


--
-- TOC entry 3129 (class 2606 OID 16747)
-- Name: person person_id; Type: CONSTRAINT; Schema: public; Owner: pedro
--

ALTER TABLE ONLY public.person
    ADD CONSTRAINT person_id PRIMARY KEY (id);


--
-- TOC entry 3130 (class 2606 OID 16748)
-- Name: person house_fk; Type: FK CONSTRAINT; Schema: public; Owner: pedro
--

ALTER TABLE ONLY public.person
    ADD CONSTRAINT house_fk FOREIGN KEY (houseid) REFERENCES public.house(id);


-- Completed on 2021-09-20 19:39:24 WEST

--
-- PostgreSQL database dump complete
--

