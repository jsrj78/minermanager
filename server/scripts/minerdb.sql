PGDMP         7            	    {           minerdb    14.6    14.4                0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false                       0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false                       0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false                       1262    41466    minerdb    DATABASE     \   CREATE DATABASE minerdb WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'en_US.UTF-8';
    DROP DATABASE minerdb;
                postgres    false            �            1255    41473    f_table_to_struct(text)    FUNCTION     �  CREATE FUNCTION public.f_table_to_struct(tb text) RETURNS text
    LANGUAGE plpgsql
    AS $$
declare
rc record;
    dbjson text;
begin
dbjson:='type '||tb||E'  struct{\r\n'||E'BaseModel \r\n';
for rc in (SELECT col_description(a.attrelid,a.attnum) as comment,a.atttypid as type,a.attname as name, a.attnotnull as notnull  
FROM pg_class as c,pg_attribute as a where c.relname = tb and a.attrelid = c.oid and a.attnum>0)
loop
raise notice '%',dbjson;
--raise notice '%',rc.type;
    if rc.name!='id' and rc.name!='createdate' and rc.name!='createuser' and rc.name!='updatedate' and rc.name!='updateuser' and rc.name!='deletedate' and rc.name!='deleteuser' and rc.name!='isdelete' then
        case
            when rc.type=2950 then
            dbjson:=dbjson||initcap(rc.name)||' uuid.UUID  `json:"'||rc.name||'" gorm:"column:'||rc.name||'"` //'|| COALESCE(rc.comment,'')||E'\r\n';
            when rc.type=1114 or rc.type=1184 then
            dbjson:=dbjson||initcap(rc.name)||' time.Time  `json:"'||rc.name||'" gorm:"column:'||rc.name||'"` //'||COALESCE(rc.comment,'')||E'\r\n';
            when rc.type=16 then
            dbjson:=dbjson||initcap(rc.name)||' bool  `json:"'||rc.name||'" gorm:"column:'||rc.name||'"` //'||COALESCE(rc.comment,'')||E'\r\n';
            when rc.type=1043 or rc.type=25 then
            dbjson:=dbjson||initcap(rc.name)||' string  `json:"'||rc.name||'" gorm:"column:'||rc.name||'"` //'||COALESCE(rc.comment,'')||E'\r\n';
when rc.type=21 or rc.type=23  then
              dbjson:=dbjson||initcap(rc.name)||' int  `json:"'||rc.name||'" gorm:"column:'||rc.name||'"` //'||COALESCE(rc.comment,'')||E'\r\n' ;      
            when rc.type=700 or rc.type=790 or rc.type=1700 then
              dbjson:=dbjson||initcap(rc.name)||' float64  `json:"'||rc.name||'" gorm:"column:'||rc.name||'"` //'||COALESCE(rc.comment,'')||E'\r\n' ;      
             when rc.type=114 then
              dbjson:=dbjson||initcap(rc.name)||' *time.Time  `json:"'||rc.name||'" gorm:"column:'||rc.name||'"` //'||COALESCE(rc.comment,'')||E'\r\n' ;      
         
else
           
            end case;
        end if;
       
end loop;
   
--raise notice '%',dbjson;

    dbjson:=dbjson||'}  //'||tb;
   
return dbjson;

end;
$$;
 1   DROP FUNCTION public.f_table_to_struct(tb text);
       public          postgres    false            �            1259    41476    t_miner    TABLE     �  CREATE TABLE public.t_miner (
    id uuid NOT NULL,
    create_date timestamp with time zone,
    create_user uuid,
    is_delete boolean,
    delete_date timestamp with time zone,
    delete_user uuid,
    update_date timestamp with time zone,
    update_user uuid,
    ip character varying(15),
    status character varying(20),
    brand character varying(20),
    username character varying(20),
    password character varying(20)
);
    DROP TABLE public.t_miner;
       public         heap    postgres    false                       0    0    COLUMN t_miner.ip    COMMENT     3   COMMENT ON COLUMN public.t_miner.ip IS 'IP地址';
          public          postgres    false    211                       0    0    COLUMN t_miner.status    COMMENT     5   COMMENT ON COLUMN public.t_miner.status IS '状态';
          public          postgres    false    211            	           0    0    COLUMN t_miner.brand    COMMENT     :   COMMENT ON COLUMN public.t_miner.brand IS '矿机品牌';
          public          postgres    false    211            
           0    0    COLUMN t_miner.username    COMMENT     :   COMMENT ON COLUMN public.t_miner.username IS '用户名';
          public          postgres    false    211                       0    0    COLUMN t_miner.password    COMMENT     7   COMMENT ON COLUMN public.t_miner.password IS '密码';
          public          postgres    false    211            �            1259    41467    t_model    TABLE       CREATE TABLE public.t_model (
    id uuid,
    create_date timestamp with time zone,
    create_user uuid,
    is_delete boolean,
    delete_date timestamp with time zone,
    delete_user uuid,
    update_date timestamp with time zone,
    update_user uuid
);
    DROP TABLE public.t_model;
       public         heap    postgres    false            �            1259    41470    t_user    TABLE     V  CREATE TABLE public.t_user (
    id uuid NOT NULL,
    create_date timestamp with time zone,
    create_user uuid,
    is_delete boolean,
    delete_date timestamp with time zone,
    delete_user uuid,
    update_date timestamp with time zone,
    update_user uuid,
    user_name character varying(50),
    password character varying(256)
);
    DROP TABLE public.t_user;
       public         heap    postgres    false                       0    0    COLUMN t_user.user_name    COMMENT     :   COMMENT ON COLUMN public.t_user.user_name IS '用户名';
          public          postgres    false    210                       0    0    COLUMN t_user.password    COMMENT     6   COMMENT ON COLUMN public.t_user.password IS '密码';
          public          postgres    false    210                       0    41476    t_miner 
   TABLE DATA           �   COPY public.t_miner (id, create_date, create_user, is_delete, delete_date, delete_user, update_date, update_user, ip, status, brand, username, password) FROM stdin;
    public          postgres    false    211   �       �          0    41467    t_model 
   TABLE DATA           ~   COPY public.t_model (id, create_date, create_user, is_delete, delete_date, delete_user, update_date, update_user) FROM stdin;
    public          postgres    false    209   �       �          0    41470    t_user 
   TABLE DATA           �   COPY public.t_user (id, create_date, create_user, is_delete, delete_date, delete_user, update_date, update_user, user_name, password) FROM stdin;
    public          postgres    false    210   �       r           2606    41480    t_miner t_miner_pkey 
   CONSTRAINT     R   ALTER TABLE ONLY public.t_miner
    ADD CONSTRAINT t_miner_pkey PRIMARY KEY (id);
 >   ALTER TABLE ONLY public.t_miner DROP CONSTRAINT t_miner_pkey;
       public            postgres    false    211            p           2606    41475    t_user t_user_pkey 
   CONSTRAINT     P   ALTER TABLE ONLY public.t_user
    ADD CONSTRAINT t_user_pkey PRIMARY KEY (id);
 <   ALTER TABLE ONLY public.t_user DROP CONSTRAINT t_user_pkey;
       public            postgres    false    210                �   x��һ�@�z�&;���
-�ix,�B��/Zhã�d�)O���
%[�jU�T�u�V���񦩱
<xŊ�� N��lh۾���q�,{b�0��|��[fNY� ��U�>b8A����׶�����RDE��ø�N�)�e7m#��,�_b)����g�q�qTF��;3C@돟��$I���5      �      x������ � �      �   ]   x�K2J�4INK�57J5�52JM�M�43�5021JL642000��Â��9�-,��PL04�b�!v�J�8ML͈1�#2J9���b���� 6�1)     