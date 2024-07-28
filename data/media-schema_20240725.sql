--
-- PostgreSQL database cluster dump
--

-- Started on 2024-07-26 03:44:41 UTC

SET default_transaction_read_only = off;

SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;

--
-- Roles
--

CREATE ROLE postgres;
ALTER ROLE postgres WITH SUPERUSER INHERIT CREATEROLE CREATEDB LOGIN REPLICATION BYPASSRLS PASSWORD 'md55305adaac499dbbc6865a44e4aa5d8b4';






--
-- Databases
--

--
-- Database "template1" dump
--

\connect template1

--
-- PostgreSQL database dump
--

-- Dumped from database version 12.19 (Debian 12.19-1.pgdg120+1)
-- Dumped by pg_dump version 12.19

-- Started on 2024-07-26 03:44:41 UTC

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

-- Completed on 2024-07-26 03:44:41 UTC

--
-- PostgreSQL database dump complete
--

--
-- Database "postgres" dump
--

\connect postgres

--
-- PostgreSQL database dump
--

-- Dumped from database version 12.19 (Debian 12.19-1.pgdg120+1)
-- Dumped by pg_dump version 12.19

-- Started on 2024-07-26 03:44:41 UTC

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
-- TOC entry 7 (class 2615 OID 16384)
-- Name: media; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA media;


ALTER SCHEMA media OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 203 (class 1259 OID 16385)
-- Name: cast; Type: TABLE; Schema: media; Owner: postgres
--

CREATE TABLE media."cast" (
    id integer NOT NULL,
    name text NOT NULL,
    birthday date,
    "desc" text,
    pic_path text
);


ALTER TABLE media."cast" OWNER TO postgres;

--
-- TOC entry 204 (class 1259 OID 16391)
-- Name: cast_id_seq; Type: SEQUENCE; Schema: media; Owner: postgres
--

CREATE SEQUENCE media.cast_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE media.cast_id_seq OWNER TO postgres;

--
-- TOC entry 3075 (class 0 OID 0)
-- Dependencies: 204
-- Name: cast_id_seq; Type: SEQUENCE OWNED BY; Schema: media; Owner: postgres
--

ALTER SEQUENCE media.cast_id_seq OWNED BY media."cast".id;


--
-- TOC entry 205 (class 1259 OID 16393)
-- Name: collection; Type: TABLE; Schema: media; Owner: postgres
--

CREATE TABLE media.collection (
    id integer NOT NULL,
    path text NOT NULL,
    disp_name text,
    parent integer,
    created timestamp without time zone DEFAULT now() NOT NULL,
    "desc" text
);


ALTER TABLE media.collection OWNER TO postgres;

--
-- TOC entry 206 (class 1259 OID 16400)
-- Name: collection_id_seq; Type: SEQUENCE; Schema: media; Owner: postgres
--

CREATE SEQUENCE media.collection_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE media.collection_id_seq OWNER TO postgres;

--
-- TOC entry 3076 (class 0 OID 0)
-- Dependencies: 206
-- Name: collection_id_seq; Type: SEQUENCE OWNED BY; Schema: media; Owner: postgres
--

ALTER SEQUENCE media.collection_id_seq OWNED BY media.collection.id;


--
-- TOC entry 207 (class 1259 OID 16402)
-- Name: collection_parent_seq; Type: SEQUENCE; Schema: media; Owner: postgres
--

CREATE SEQUENCE media.collection_parent_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE media.collection_parent_seq OWNER TO postgres;

--
-- TOC entry 3077 (class 0 OID 0)
-- Dependencies: 207
-- Name: collection_parent_seq; Type: SEQUENCE OWNED BY; Schema: media; Owner: postgres
--

ALTER SEQUENCE media.collection_parent_seq OWNED BY media.collection.parent;


--
-- TOC entry 208 (class 1259 OID 16404)
-- Name: entry; Type: TABLE; Schema: media; Owner: postgres
--

CREATE TABLE media.entry (
    id integer NOT NULL,
    path text NOT NULL,
    disp_name text,
    "desc" text,
    parent integer,
    thumb_static text,
    thumb_dynamic text[],
    created timestamp without time zone DEFAULT now() NOT NULL,
    updated timestamp without time zone DEFAULT now() NOT NULL,
    aired timestamp without time zone,
    content_files text[]
);


ALTER TABLE media.entry OWNER TO postgres;

--
-- TOC entry 209 (class 1259 OID 16412)
-- Name: entry_cast; Type: TABLE; Schema: media; Owner: postgres
--

CREATE TABLE media.entry_cast (
    cast_id integer NOT NULL,
    entry_id integer NOT NULL
);


ALTER TABLE media.entry_cast OWNER TO postgres;

--
-- TOC entry 210 (class 1259 OID 16415)
-- Name: entry_cast_cast_id_seq; Type: SEQUENCE; Schema: media; Owner: postgres
--

CREATE SEQUENCE media.entry_cast_cast_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE media.entry_cast_cast_id_seq OWNER TO postgres;

--
-- TOC entry 3078 (class 0 OID 0)
-- Dependencies: 210
-- Name: entry_cast_cast_id_seq; Type: SEQUENCE OWNED BY; Schema: media; Owner: postgres
--

ALTER SEQUENCE media.entry_cast_cast_id_seq OWNED BY media.entry_cast.cast_id;


--
-- TOC entry 211 (class 1259 OID 16417)
-- Name: entry_cast_entry_id_seq; Type: SEQUENCE; Schema: media; Owner: postgres
--

CREATE SEQUENCE media.entry_cast_entry_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE media.entry_cast_entry_id_seq OWNER TO postgres;

--
-- TOC entry 3079 (class 0 OID 0)
-- Dependencies: 211
-- Name: entry_cast_entry_id_seq; Type: SEQUENCE OWNED BY; Schema: media; Owner: postgres
--

ALTER SEQUENCE media.entry_cast_entry_id_seq OWNED BY media.entry_cast.entry_id;


--
-- TOC entry 212 (class 1259 OID 16419)
-- Name: entry_id_seq; Type: SEQUENCE; Schema: media; Owner: postgres
--

CREATE SEQUENCE media.entry_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE media.entry_id_seq OWNER TO postgres;

--
-- TOC entry 3080 (class 0 OID 0)
-- Dependencies: 212
-- Name: entry_id_seq; Type: SEQUENCE OWNED BY; Schema: media; Owner: postgres
--

ALTER SEQUENCE media.entry_id_seq OWNED BY media.entry.id;


--
-- TOC entry 213 (class 1259 OID 16421)
-- Name: entry_parent_seq; Type: SEQUENCE; Schema: media; Owner: postgres
--

CREATE SEQUENCE media.entry_parent_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE media.entry_parent_seq OWNER TO postgres;

--
-- TOC entry 3081 (class 0 OID 0)
-- Dependencies: 213
-- Name: entry_parent_seq; Type: SEQUENCE OWNED BY; Schema: media; Owner: postgres
--

ALTER SEQUENCE media.entry_parent_seq OWNED BY media.entry.parent;


--
-- TOC entry 214 (class 1259 OID 16423)
-- Name: entry_tag; Type: TABLE; Schema: media; Owner: postgres
--

CREATE TABLE media.entry_tag (
    entry_id integer NOT NULL,
    tag_id integer NOT NULL
);


ALTER TABLE media.entry_tag OWNER TO postgres;

--
-- TOC entry 215 (class 1259 OID 16426)
-- Name: entry_tag_entry_id_seq; Type: SEQUENCE; Schema: media; Owner: postgres
--

CREATE SEQUENCE media.entry_tag_entry_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE media.entry_tag_entry_id_seq OWNER TO postgres;

--
-- TOC entry 3082 (class 0 OID 0)
-- Dependencies: 215
-- Name: entry_tag_entry_id_seq; Type: SEQUENCE OWNED BY; Schema: media; Owner: postgres
--

ALTER SEQUENCE media.entry_tag_entry_id_seq OWNED BY media.entry_tag.entry_id;


--
-- TOC entry 216 (class 1259 OID 16428)
-- Name: entry_tag_tag_id_seq; Type: SEQUENCE; Schema: media; Owner: postgres
--

CREATE SEQUENCE media.entry_tag_tag_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE media.entry_tag_tag_id_seq OWNER TO postgres;

--
-- TOC entry 3083 (class 0 OID 0)
-- Dependencies: 216
-- Name: entry_tag_tag_id_seq; Type: SEQUENCE OWNED BY; Schema: media; Owner: postgres
--

ALTER SEQUENCE media.entry_tag_tag_id_seq OWNED BY media.entry_tag.tag_id;


--
-- TOC entry 217 (class 1259 OID 16430)
-- Name: tag; Type: TABLE; Schema: media; Owner: postgres
--

CREATE TABLE media.tag (
    id integer NOT NULL,
    name text NOT NULL,
    parent integer
);


ALTER TABLE media.tag OWNER TO postgres;

--
-- TOC entry 218 (class 1259 OID 16436)
-- Name: tag_id_seq; Type: SEQUENCE; Schema: media; Owner: postgres
--

CREATE SEQUENCE media.tag_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE media.tag_id_seq OWNER TO postgres;

--
-- TOC entry 3084 (class 0 OID 0)
-- Dependencies: 218
-- Name: tag_id_seq; Type: SEQUENCE OWNED BY; Schema: media; Owner: postgres
--

ALTER SEQUENCE media.tag_id_seq OWNED BY media.tag.id;


--
-- TOC entry 219 (class 1259 OID 16438)
-- Name: tag_parent_seq; Type: SEQUENCE; Schema: media; Owner: postgres
--

CREATE SEQUENCE media.tag_parent_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE media.tag_parent_seq OWNER TO postgres;

--
-- TOC entry 3085 (class 0 OID 0)
-- Dependencies: 219
-- Name: tag_parent_seq; Type: SEQUENCE OWNED BY; Schema: media; Owner: postgres
--

ALTER SEQUENCE media.tag_parent_seq OWNED BY media.tag.parent;


--
-- TOC entry 220 (class 1259 OID 16440)
-- Name: user; Type: TABLE; Schema: media; Owner: postgres
--

CREATE TABLE media."user" (
    id integer NOT NULL,
    name text,
    email text,
    password text,
    created timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE media."user" OWNER TO postgres;

--
-- TOC entry 221 (class 1259 OID 16446)
-- Name: user_id_seq; Type: SEQUENCE; Schema: media; Owner: postgres
--

CREATE SEQUENCE media.user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE media.user_id_seq OWNER TO postgres;

--
-- TOC entry 3086 (class 0 OID 0)
-- Dependencies: 221
-- Name: user_id_seq; Type: SEQUENCE OWNED BY; Schema: media; Owner: postgres
--

ALTER SEQUENCE media.user_id_seq OWNED BY media."user".id;


--
-- TOC entry 2902 (class 2604 OID 16448)
-- Name: cast id; Type: DEFAULT; Schema: media; Owner: postgres
--

ALTER TABLE ONLY media."cast" ALTER COLUMN id SET DEFAULT nextval('media.cast_id_seq'::regclass);


--
-- TOC entry 2904 (class 2604 OID 16449)
-- Name: collection id; Type: DEFAULT; Schema: media; Owner: postgres
--

ALTER TABLE ONLY media.collection ALTER COLUMN id SET DEFAULT nextval('media.collection_id_seq'::regclass);


--
-- TOC entry 2907 (class 2604 OID 16450)
-- Name: entry id; Type: DEFAULT; Schema: media; Owner: postgres
--

ALTER TABLE ONLY media.entry ALTER COLUMN id SET DEFAULT nextval('media.entry_id_seq'::regclass);


--
-- TOC entry 2908 (class 2604 OID 16451)
-- Name: entry_tag entry_id; Type: DEFAULT; Schema: media; Owner: postgres
--

ALTER TABLE ONLY media.entry_tag ALTER COLUMN entry_id SET DEFAULT nextval('media.entry_tag_entry_id_seq'::regclass);


--
-- TOC entry 2909 (class 2604 OID 16452)
-- Name: entry_tag tag_id; Type: DEFAULT; Schema: media; Owner: postgres
--

ALTER TABLE ONLY media.entry_tag ALTER COLUMN tag_id SET DEFAULT nextval('media.entry_tag_tag_id_seq'::regclass);


--
-- TOC entry 2910 (class 2604 OID 16453)
-- Name: tag id; Type: DEFAULT; Schema: media; Owner: postgres
--

ALTER TABLE ONLY media.tag ALTER COLUMN id SET DEFAULT nextval('media.tag_id_seq'::regclass);


--
-- TOC entry 2911 (class 2604 OID 16455)
-- Name: user id; Type: DEFAULT; Schema: media; Owner: postgres
--

ALTER TABLE ONLY media."user" ALTER COLUMN id SET DEFAULT nextval('media.user_id_seq'::regclass);


--
-- TOC entry 2914 (class 2606 OID 16457)
-- Name: cast cast_pkey; Type: CONSTRAINT; Schema: media; Owner: postgres
--

ALTER TABLE ONLY media."cast"
    ADD CONSTRAINT cast_pkey PRIMARY KEY (id);


--
-- TOC entry 2916 (class 2606 OID 16459)
-- Name: collection pk_collection; Type: CONSTRAINT; Schema: media; Owner: postgres
--

ALTER TABLE ONLY media.collection
    ADD CONSTRAINT pk_collection PRIMARY KEY (id);


--
-- TOC entry 2920 (class 2606 OID 16461)
-- Name: entry pk_entry; Type: CONSTRAINT; Schema: media; Owner: postgres
--

ALTER TABLE ONLY media.entry
    ADD CONSTRAINT pk_entry PRIMARY KEY (id);


--
-- TOC entry 2924 (class 2606 OID 16463)
-- Name: entry_cast pk_entry_cast; Type: CONSTRAINT; Schema: media; Owner: postgres
--

ALTER TABLE ONLY media.entry_cast
    ADD CONSTRAINT pk_entry_cast PRIMARY KEY (entry_id, cast_id);


--
-- TOC entry 2926 (class 2606 OID 16465)
-- Name: entry_tag pk_entry_tag; Type: CONSTRAINT; Schema: media; Owner: postgres
--

ALTER TABLE ONLY media.entry_tag
    ADD CONSTRAINT pk_entry_tag PRIMARY KEY (entry_id, tag_id);


--
-- TOC entry 2928 (class 2606 OID 16467)
-- Name: tag pk_tag; Type: CONSTRAINT; Schema: media; Owner: postgres
--

ALTER TABLE ONLY media.tag
    ADD CONSTRAINT pk_tag PRIMARY KEY (id);


--
-- TOC entry 2932 (class 2606 OID 16469)
-- Name: user pk_user; Type: CONSTRAINT; Schema: media; Owner: postgres
--

ALTER TABLE ONLY media."user"
    ADD CONSTRAINT pk_user PRIMARY KEY (id);


--
-- TOC entry 2934 (class 2606 OID 16471)
-- Name: user u_email; Type: CONSTRAINT; Schema: media; Owner: postgres
--

ALTER TABLE ONLY media."user"
    ADD CONSTRAINT u_email UNIQUE (email);


--
-- TOC entry 2936 (class 2606 OID 16473)
-- Name: user u_name; Type: CONSTRAINT; Schema: media; Owner: postgres
--

ALTER TABLE ONLY media."user"
    ADD CONSTRAINT u_name UNIQUE (name);


--
-- TOC entry 2930 (class 2606 OID 16517)
-- Name: tag u_tagname; Type: CONSTRAINT; Schema: media; Owner: postgres
--

ALTER TABLE ONLY media.tag
    ADD CONSTRAINT u_tagname UNIQUE (name);


--
-- TOC entry 2922 (class 2606 OID 16475)
-- Name: entry unique_parent_path; Type: CONSTRAINT; Schema: media; Owner: postgres
--

ALTER TABLE ONLY media.entry
    ADD CONSTRAINT unique_parent_path UNIQUE (path, parent);


--
-- TOC entry 2918 (class 2606 OID 16477)
-- Name: collection unique_path; Type: CONSTRAINT; Schema: media; Owner: postgres
--

ALTER TABLE ONLY media.collection
    ADD CONSTRAINT unique_path UNIQUE (path);


--
-- TOC entry 2939 (class 2606 OID 16478)
-- Name: entry_cast fk_cast; Type: FK CONSTRAINT; Schema: media; Owner: postgres
--

ALTER TABLE ONLY media.entry_cast
    ADD CONSTRAINT fk_cast FOREIGN KEY (cast_id) REFERENCES media."cast"(id);


--
-- TOC entry 2937 (class 2606 OID 16483)
-- Name: collection fk_collection_self; Type: FK CONSTRAINT; Schema: media; Owner: postgres
--

ALTER TABLE ONLY media.collection
    ADD CONSTRAINT fk_collection_self FOREIGN KEY (parent) REFERENCES media.collection(id) NOT VALID;


--
-- TOC entry 2940 (class 2606 OID 16488)
-- Name: entry_cast fk_entry; Type: FK CONSTRAINT; Schema: media; Owner: postgres
--

ALTER TABLE ONLY media.entry_cast
    ADD CONSTRAINT fk_entry FOREIGN KEY (entry_id) REFERENCES media.entry(id);


--
-- TOC entry 2941 (class 2606 OID 16493)
-- Name: entry_tag fk_entry; Type: FK CONSTRAINT; Schema: media; Owner: postgres
--

ALTER TABLE ONLY media.entry_tag
    ADD CONSTRAINT fk_entry FOREIGN KEY (entry_id) REFERENCES media.entry(id);


--
-- TOC entry 2938 (class 2606 OID 16498)
-- Name: entry fk_entry_parent; Type: FK CONSTRAINT; Schema: media; Owner: postgres
--

ALTER TABLE ONLY media.entry
    ADD CONSTRAINT fk_entry_parent FOREIGN KEY (parent) REFERENCES media.collection(id);


--
-- TOC entry 2943 (class 2606 OID 16503)
-- Name: tag fk_tag; Type: FK CONSTRAINT; Schema: media; Owner: postgres
--

ALTER TABLE ONLY media.tag
    ADD CONSTRAINT fk_tag FOREIGN KEY (parent) REFERENCES media.tag(id) NOT VALID;


--
-- TOC entry 2942 (class 2606 OID 16508)
-- Name: entry_tag fk_tag; Type: FK CONSTRAINT; Schema: media; Owner: postgres
--

ALTER TABLE ONLY media.entry_tag
    ADD CONSTRAINT fk_tag FOREIGN KEY (tag_id) REFERENCES media.tag(id);


-- Completed on 2024-07-26 03:44:41 UTC

--
-- PostgreSQL database dump complete
--

-- Completed on 2024-07-26 03:44:41 UTC

--
-- PostgreSQL database cluster dump complete
--

