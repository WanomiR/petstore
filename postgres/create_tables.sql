--
-- PostgreSQL database dump
--

-- Dumped from database version 16.3
-- Dumped by pg_dump version 16.3

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: categories; Type: TABLE; Schema: public; Owner: user
--

CREATE TABLE public.categories (
    id integer NOT NULL,
    name character varying(255)
);


ALTER TABLE public.categories OWNER TO "user";

--
-- Name: categories_id_seq; Type: SEQUENCE; Schema: public; Owner: user
--

CREATE SEQUENCE public.categories_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.categories_id_seq OWNER TO "user";

--
-- Name: categories_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: user
--

ALTER SEQUENCE public.categories_id_seq OWNED BY public.categories.id;


--
-- Name: pet_tags; Type: TABLE; Schema: public; Owner: user
--

CREATE TABLE public.pet_tags (
    id integer NOT NULL,
    pet_id integer,
    tag_id integer
);


ALTER TABLE public.pet_tags OWNER TO "user";

--
-- Name: pet_tags_id_seq; Type: SEQUENCE; Schema: public; Owner: user
--

CREATE SEQUENCE public.pet_tags_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.pet_tags_id_seq OWNER TO "user";

--
-- Name: pet_tags_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: user
--

ALTER SEQUENCE public.pet_tags_id_seq OWNED BY public.pet_tags.id;


--
-- Name: pets; Type: TABLE; Schema: public; Owner: user
--

CREATE TABLE public.pets (
    id integer NOT NULL,
    category_id integer,
    name character varying(255),
    status character varying(255),
    is_deleted boolean,
    CONSTRAINT check_status CHECK (((status)::text = ANY ((ARRAY['available'::character varying, 'pending'::character varying, 'sold'::character varying])::text[])))
);


ALTER TABLE public.pets OWNER TO "user";

--
-- Name: pets_id_seq; Type: SEQUENCE; Schema: public; Owner: user
--

CREATE SEQUENCE public.pets_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.pets_id_seq OWNER TO "user";

--
-- Name: pets_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: user
--

ALTER SEQUENCE public.pets_id_seq OWNED BY public.pets.id;


--
-- Name: photo_urls; Type: TABLE; Schema: public; Owner: user
--

CREATE TABLE public.photo_urls (
    id integer NOT NULL,
    pet_id integer,
    url character varying(511)
);


ALTER TABLE public.photo_urls OWNER TO "user";

--
-- Name: photo_urls_id_seq; Type: SEQUENCE; Schema: public; Owner: user
--

CREATE SEQUENCE public.photo_urls_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.photo_urls_id_seq OWNER TO "user";

--
-- Name: photo_urls_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: user
--

ALTER SEQUENCE public.photo_urls_id_seq OWNED BY public.photo_urls.id;


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: user
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations OWNER TO "user";

--
-- Name: store; Type: TABLE; Schema: public; Owner: user
--

CREATE TABLE public.store (
    id integer NOT NULL,
    pet_id integer,
    quantity integer,
    ship_date timestamp without time zone,
    status character varying(255),
    is_complete boolean,
    CONSTRAINT check_status CHECK (((status)::text = ANY ((ARRAY['placed'::character varying, 'approved'::character varying, 'delivered'::character varying])::text[])))
);


ALTER TABLE public.store OWNER TO "user";

--
-- Name: store_id_seq; Type: SEQUENCE; Schema: public; Owner: user
--

CREATE SEQUENCE public.store_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.store_id_seq OWNER TO "user";

--
-- Name: store_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: user
--

ALTER SEQUENCE public.store_id_seq OWNED BY public.store.id;


--
-- Name: tags; Type: TABLE; Schema: public; Owner: user
--

CREATE TABLE public.tags (
    id integer NOT NULL,
    name character varying(255)
);


ALTER TABLE public.tags OWNER TO "user";

--
-- Name: tags_id_seq; Type: SEQUENCE; Schema: public; Owner: user
--

CREATE SEQUENCE public.tags_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.tags_id_seq OWNER TO "user";

--
-- Name: tags_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: user
--

ALTER SEQUENCE public.tags_id_seq OWNED BY public.tags.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: user
--

CREATE TABLE public.users (
    id integer NOT NULL,
    username character varying(255),
    first_name character varying(255),
    last_name character varying(255),
    email character varying(255),
    password character varying(255),
    phone character varying(255),
    user_status integer
);


ALTER TABLE public.users OWNER TO "user";

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: user
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO "user";

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: user
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: categories id; Type: DEFAULT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.categories ALTER COLUMN id SET DEFAULT nextval('public.categories_id_seq'::regclass);


--
-- Name: pet_tags id; Type: DEFAULT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.pet_tags ALTER COLUMN id SET DEFAULT nextval('public.pet_tags_id_seq'::regclass);


--
-- Name: pets id; Type: DEFAULT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.pets ALTER COLUMN id SET DEFAULT nextval('public.pets_id_seq'::regclass);


--
-- Name: photo_urls id; Type: DEFAULT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.photo_urls ALTER COLUMN id SET DEFAULT nextval('public.photo_urls_id_seq'::regclass);


--
-- Name: store id; Type: DEFAULT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.store ALTER COLUMN id SET DEFAULT nextval('public.store_id_seq'::regclass);


--
-- Name: tags id; Type: DEFAULT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.tags ALTER COLUMN id SET DEFAULT nextval('public.tags_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: categories; Type: TABLE DATA; Schema: public; Owner: user
--

COPY public.categories (id, name) FROM stdin;
1	cat
2	dog
3	rodent
\.


--
-- Data for Name: pet_tags; Type: TABLE DATA; Schema: public; Owner: user
--

COPY public.pet_tags (id, pet_id, tag_id) FROM stdin;
1	1	1
2	1	7
3	2	4
4	2	6
5	3	1
6	4	5
7	5	1
8	5	2
9	5	7
10	6	2
11	6	6
12	7	1
13	8	3
14	9	4
\.


--
-- Data for Name: pets; Type: TABLE DATA; Schema: public; Owner: user
--

COPY public.pets (id, category_id, name, status, is_deleted) FROM stdin;
1	1	Poppy	available	f
2	1	Bella	pending	f
3	1	Tilly	sold	f
4	2	Abby	available	f
5	2	Bailey	pending	f
6	2	Rex	sold	f
7	3	Basil	available	f
8	3	Danger Mouse	pending	f
9	3	Jerry	sold	f
\.


--
-- Data for Name: photo_urls; Type: TABLE DATA; Schema: public; Owner: user
--

COPY public.photo_urls (id, pet_id, url) FROM stdin;
1	1	https://wallbox.ru/wallpapers/main/201152/koshki-392426de15fb.jpg
2	1	https://wallbox.ru/resize/1024x1024/wallpapers/main/201634/8b7e73ae5927008.jpg
3	2	https://pixy.org/src/471/4710119.jpg
4	3	https://jooinn.com/images/happy-cat-resting-6.jpg
5	4	https://wallpapers.com/images/hd/dog-pictures-os09dhwexb80d990.jpg
6	5	https://c.pxhere.com/photos/f6/1d/dog_pet_small_dog-912658.jpg!d
7	6	https://jooinn.com/images/pet-dog-142.jpg
8	7	https://i.pinimg.com/originals/59/df/fb/59dffb52f7435ce31979f7e03ce02ab4.jpg
9	8	https://i.pinimg.com/originals/fe/e8/8a/fee88a1d551c31b2217d999146bfdeb1.jpg
10	9	https://i.pinimg.com/originals/3a/69/ae/3a69aee66a3f324915ee3085baf9c6c4.jpg
\.


--
-- Data for Name: schema_migrations; Type: TABLE DATA; Schema: public; Owner: user
--

COPY public.schema_migrations (version, dirty) FROM stdin;
1	f
\.


--
-- Data for Name: store; Type: TABLE DATA; Schema: public; Owner: user
--

COPY public.store (id, pet_id, quantity, ship_date, status, is_complete) FROM stdin;
1	2	1	2024-07-31 12:11:14.265084	placed	f
2	3	1	2024-07-31 12:11:14.265084	delivered	t
3	5	1	2024-07-31 12:11:14.265084	approved	f
4	6	1	2024-07-31 12:11:14.265084	delivered	t
5	8	1	2024-07-31 12:11:14.265084	placed	f
6	9	1	2024-07-31 12:11:14.265084	delivered	t
\.


--
-- Data for Name: tags; Type: TABLE DATA; Schema: public; Owner: user
--

COPY public.tags (id, name) FROM stdin;
1	fluffy
2	funny
3	kind
4	playful
5	calm
6	happy
7	energetic
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: user
--

COPY public.users (id, username, first_name, last_name, email, password, phone, user_status) FROM stdin;
1	wanomir	Ivan	Romadin	wanomir@yandex.ru	$2a$10$TzogjOIjVZ9fY8/J.1EgOOlV9E1IOSGTC5WWYoP.tDewfMkYUAUXu	7-999-999-99-99	0
2	johndoe001	John	Doe	john.doe@gmail.com	$2a$10$TzogjOIjVZ9fY8/J.1EgOOlV9E1IOSGTC5WWYoP.tDewfMkYUAUXu	7-999-999-99-99	0
3	dragonrider	Rhaenyra	Targaryen	r.targaryen@dragonstone.com	$2a$10$TzogjOIjVZ9fY8/J.1EgOOlV9E1IOSGTC5WWYoP.tDewfMkYUAUXu	7-999-999-99-99	0
\.


--
-- Name: categories_id_seq; Type: SEQUENCE SET; Schema: public; Owner: user
--

SELECT pg_catalog.setval('public.categories_id_seq', 3, true);


--
-- Name: pet_tags_id_seq; Type: SEQUENCE SET; Schema: public; Owner: user
--

SELECT pg_catalog.setval('public.pet_tags_id_seq', 14, true);


--
-- Name: pets_id_seq; Type: SEQUENCE SET; Schema: public; Owner: user
--

SELECT pg_catalog.setval('public.pets_id_seq', 9, true);


--
-- Name: photo_urls_id_seq; Type: SEQUENCE SET; Schema: public; Owner: user
--

SELECT pg_catalog.setval('public.photo_urls_id_seq', 10, true);


--
-- Name: store_id_seq; Type: SEQUENCE SET; Schema: public; Owner: user
--

SELECT pg_catalog.setval('public.store_id_seq', 6, true);


--
-- Name: tags_id_seq; Type: SEQUENCE SET; Schema: public; Owner: user
--

SELECT pg_catalog.setval('public.tags_id_seq', 7, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: user
--

SELECT pg_catalog.setval('public.users_id_seq', 3, true);


--
-- Name: categories categories_pkey; Type: CONSTRAINT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_pkey PRIMARY KEY (id);


--
-- Name: pet_tags pet_tags_pkey; Type: CONSTRAINT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.pet_tags
    ADD CONSTRAINT pet_tags_pkey PRIMARY KEY (id);


--
-- Name: pets pets_pkey; Type: CONSTRAINT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.pets
    ADD CONSTRAINT pets_pkey PRIMARY KEY (id);


--
-- Name: photo_urls photo_urls_pkey; Type: CONSTRAINT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.photo_urls
    ADD CONSTRAINT photo_urls_pkey PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: store store_pkey; Type: CONSTRAINT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.store
    ADD CONSTRAINT store_pkey PRIMARY KEY (id);


--
-- Name: tags tags_pkey; Type: CONSTRAINT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.tags
    ADD CONSTRAINT tags_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: users users_username_key; Type: CONSTRAINT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_key UNIQUE (username);


--
-- PostgreSQL database dump complete
--

