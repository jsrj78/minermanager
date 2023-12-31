PGDMP         )            
    {            minerdb    15.4    15.4 P               0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false                       0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false                       0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false                       1262    24576    minerdb    DATABASE     �   CREATE DATABASE minerdb WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'English_United States.1252';
    DROP DATABASE minerdb;
                postgres    false            �            1255    24577    f_table_to_struct(text)    FUNCTION     �  CREATE FUNCTION public.f_table_to_struct(tb text) RETURNS text
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
       public          postgres    false            �            1259    24598    t_boxes    TABLE       CREATE TABLE public.t_boxes (
    id uuid NOT NULL,
    create_date timestamp with time zone,
    create_user uuid,
    is_delete boolean,
    delete_date timestamp with time zone,
    delete_user uuid,
    update_date timestamp with time zone,
    update_user uuid,
    lineid uuid,
    total_miners integer,
    offline_miners integer,
    normal_miners integer,
    empty_miners integer,
    total_place integer,
    ip_start character varying(200),
    ip_end character varying(200),
    islift boolean,
    empty_place integer
);
    DROP TABLE public.t_boxes;
       public         heap    postgres    false                       0    0    TABLE t_boxes    COMMENT     K   COMMENT ON TABLE public.t_boxes IS '货架或仓库中的每个小盒子';
          public          postgres    false    218                        0    0    COLUMN t_boxes.lineid    COMMENT     =   COMMENT ON COLUMN public.t_boxes.lineid IS '所在货架id';
          public          postgres    false    218            !           0    0    COLUMN t_boxes.total_miners    COMMENT     D   COMMENT ON COLUMN public.t_boxes.total_miners IS '矿机总数量';
          public          postgres    false    218            "           0    0    COLUMN t_boxes.offline_miners    COMMENT     I   COMMENT ON COLUMN public.t_boxes.offline_miners IS '离线矿机数量';
          public          postgres    false    218            #           0    0    COLUMN t_boxes.normal_miners    COMMENT     H   COMMENT ON COLUMN public.t_boxes.normal_miners IS '正常在线数量';
          public          postgres    false    218            $           0    0    COLUMN t_boxes.empty_miners    COMMENT     G   COMMENT ON COLUMN public.t_boxes.empty_miners IS '空壳机器数量';
          public          postgres    false    218            %           0    0    COLUMN t_boxes.total_place    COMMENT     C   COMMENT ON COLUMN public.t_boxes.total_place IS '总机位数量';
          public          postgres    false    218            &           0    0    COLUMN t_boxes.islift    COMMENT     D   COMMENT ON COLUMN public.t_boxes.islift IS '是否需要升降机';
          public          postgres    false    218            '           0    0    COLUMN t_boxes.empty_place    COMMENT     C   COMMENT ON COLUMN public.t_boxes.empty_place IS '空机位数量';
          public          postgres    false    218            �            1259    24603    t_brand    TABLE     �  CREATE TABLE public.t_brand (
    id uuid NOT NULL,
    create_date timestamp with time zone,
    create_user uuid,
    is_delete boolean,
    delete_date timestamp with time zone,
    delete_user uuid,
    update_date timestamp with time zone,
    update_user uuid,
    code character varying(200),
    title character varying(200),
    iorder integer,
    password character varying,
    username character varying
);
    DROP TABLE public.t_brand;
       public         heap    postgres    false            (           0    0    TABLE t_brand    COMMENT     B   COMMENT ON TABLE public.t_brand IS '机器品牌或生产厂家';
          public          postgres    false    219            )           0    0    COLUMN t_brand.code    COMMENT     9   COMMENT ON COLUMN public.t_brand.code IS '厂家编号';
          public          postgres    false    219            *           0    0    COLUMN t_brand.title    COMMENT     :   COMMENT ON COLUMN public.t_brand.title IS '厂家名称';
          public          postgres    false    219            +           0    0    COLUMN t_brand.iorder    COMMENT     5   COMMENT ON COLUMN public.t_brand.iorder IS '序号';
          public          postgres    false    219            ,           0    0    COLUMN t_brand.password    COMMENT     @   COMMENT ON COLUMN public.t_brand.password IS '默认用户名';
          public          postgres    false    219            -           0    0    COLUMN t_brand.username    COMMENT     @   COMMENT ON COLUMN public.t_brand.username IS '默认用户名';
          public          postgres    false    219            �            1259    24591    t_line    TABLE     T  CREATE TABLE public.t_line (
    id uuid NOT NULL,
    create_date timestamp with time zone,
    create_user uuid,
    is_delete boolean,
    delete_date timestamp with time zone,
    delete_user uuid,
    update_date timestamp with time zone,
    update_user uuid,
    code character varying(200),
    title character varying(500),
    boxes integer,
    ip character varying(20),
    rule_floors integer,
    rule_miners integer,
    rule_box_columns integer,
    actual_total_miners integer,
    actual_empty_place character varying,
    actual_empty_miners integer,
    actual_normal_miners integer,
    actual_offline_miners integer,
    islift boolean,
    standfloor smallint,
    ip_start integer,
    miner_brandid character varying,
    rule_username character varying(100),
    rule_password character varying(100),
    box_begin integer
);
    DROP TABLE public.t_line;
       public         heap    postgres    false            .           0    0    TABLE t_line    COMMENT     V   COMMENT ON TABLE public.t_line IS '机器存放的货架或仓库，我们称作线';
          public          postgres    false    217            /           0    0    COLUMN t_line.code    COMMENT     5   COMMENT ON COLUMN public.t_line.code IS '线编号';
          public          postgres    false    217            0           0    0    COLUMN t_line.title    COMMENT     6   COMMENT ON COLUMN public.t_line.title IS '线名称';
          public          postgres    false    217            1           0    0    COLUMN t_line.boxes    COMMENT     ?   COMMENT ON COLUMN public.t_line.boxes IS '线的总box数量';
          public          postgres    false    217            2           0    0    COLUMN t_line.ip    COMMENT     B   COMMENT ON COLUMN public.t_line.ip IS 'ip地址（开始2位）';
          public          postgres    false    217            3           0    0    COLUMN t_line.rule_floors    COMMENT     H   COMMENT ON COLUMN public.t_line.rule_floors IS '每个盒子多少层';
          public          postgres    false    217            4           0    0    COLUMN t_line.rule_miners    COMMENT     H   COMMENT ON COLUMN public.t_line.rule_miners IS '每层多少个miners';
          public          postgres    false    217            5           0    0    COLUMN t_line.rule_box_columns    COMMENT     m   COMMENT ON COLUMN public.t_line.rule_box_columns IS '每个盒子包括多少列，默认1列，有的3列';
          public          postgres    false    217            6           0    0 !   COLUMN t_line.actual_total_miners    COMMENT     P   COMMENT ON COLUMN public.t_line.actual_total_miners IS '实际总机器数量';
          public          postgres    false    217            7           0    0     COLUMN t_line.actual_empty_place    COMMENT     L   COMMENT ON COLUMN public.t_line.actual_empty_place IS '实际的空机位';
          public          postgres    false    217            8           0    0 !   COLUMN t_line.actual_empty_miners    COMMENT     S   COMMENT ON COLUMN public.t_line.actual_empty_miners IS '其中，空机器数量';
          public          postgres    false    217            9           0    0 "   COLUMN t_line.actual_normal_miners    COMMENT     Q   COMMENT ON COLUMN public.t_line.actual_normal_miners IS '所有正常的机器';
          public          postgres    false    217            :           0    0 #   COLUMN t_line.actual_offline_miners    COMMENT     R   COMMENT ON COLUMN public.t_line.actual_offline_miners IS '所有离线的机器';
          public          postgres    false    217            ;           0    0    COLUMN t_line.islift    COMMENT     O   COMMENT ON COLUMN public.t_line.islift IS '盒子高度是否需要升降机';
          public          postgres    false    217            <           0    0    COLUMN t_line.standfloor    COMMENT     h   COMMENT ON COLUMN public.t_line.standfloor IS '如果用升降机，正常人站立能够到的层数';
          public          postgres    false    217            =           0    0    COLUMN t_line.ip_start    COMMENT     n   COMMENT ON COLUMN public.t_line.ip_start IS 'ip从第几位开始，默认从1，也有从6或者25的     ';
          public          postgres    false    217            >           0    0    COLUMN t_line.miner_brandid    COMMENT     A   COMMENT ON COLUMN public.t_line.miner_brandid IS '机器品牌';
          public          postgres    false    217            ?           0    0    COLUMN t_line.rule_username    COMMENT     M   COMMENT ON COLUMN public.t_line.rule_username IS '机器实际的用户名';
          public          postgres    false    217            @           0    0    COLUMN t_line.rule_password    COMMENT     J   COMMENT ON COLUMN public.t_line.rule_password IS '机器实际的密码';
          public          postgres    false    217            A           0    0    COLUMN t_line.box_begin    COMMENT     C   COMMENT ON COLUMN public.t_line.box_begin IS '盒子开始编号';
          public          postgres    false    217            �            1259    24578    t_miner    TABLE     M  CREATE TABLE public.t_miner (
    id uuid NOT NULL,
    create_date timestamp with time zone,
    create_user uuid,
    is_delete boolean,
    delete_date timestamp with time zone,
    delete_user uuid,
    update_date timestamp with time zone,
    update_user uuid,
    status character varying(20),
    brand character varying(20),
    username character varying(20),
    password character varying(20),
    boxid uuid,
    default_username character varying(100),
    default_password character varying(100),
    ip inet,
    cols smallint,
    rows smallint,
    child_box smallint
);
    DROP TABLE public.t_miner;
       public         heap    postgres    false            B           0    0    TABLE t_miner    COMMENT     6   COMMENT ON TABLE public.t_miner IS '具体的机器';
          public          postgres    false    214            C           0    0    COLUMN t_miner.status    COMMENT     5   COMMENT ON COLUMN public.t_miner.status IS '状态';
          public          postgres    false    214            D           0    0    COLUMN t_miner.brand    COMMENT     :   COMMENT ON COLUMN public.t_miner.brand IS '矿机品牌';
          public          postgres    false    214            E           0    0    COLUMN t_miner.username    COMMENT     :   COMMENT ON COLUMN public.t_miner.username IS '用户名';
          public          postgres    false    214            F           0    0    COLUMN t_miner.password    COMMENT     7   COMMENT ON COLUMN public.t_miner.password IS '密码';
          public          postgres    false    214            G           0    0    COLUMN t_miner.boxid    COMMENT     ?   COMMENT ON COLUMN public.t_miner.boxid IS '所在盒子的id';
          public          postgres    false    214            H           0    0    COLUMN t_miner.default_username    COMMENT     H   COMMENT ON COLUMN public.t_miner.default_username IS '默认用户名';
          public          postgres    false    214            I           0    0    COLUMN t_miner.default_password    COMMENT     E   COMMENT ON COLUMN public.t_miner.default_password IS '默认密码';
          public          postgres    false    214            J           0    0    COLUMN t_miner.ip    COMMENT     3   COMMENT ON COLUMN public.t_miner.ip IS 'ip地址';
          public          postgres    false    214            K           0    0    COLUMN t_miner.cols    COMMENT     O   COMMENT ON COLUMN public.t_miner.cols IS '所在列，大部分机器默认1';
          public          postgres    false    214            L           0    0    COLUMN t_miner.rows    COMMENT     6   COMMENT ON COLUMN public.t_miner.rows IS '所在行';
          public          postgres    false    214            M           0    0    COLUMN t_miner.child_box    COMMENT     c   COMMENT ON COLUMN public.t_miner.child_box IS '在盒子里的第一个小盒子，一般默认1';
          public          postgres    false    214            �            1259    24581    t_model    TABLE     �  CREATE TABLE public.t_model (
    id uuid NOT NULL,
    create_date timestamp with time zone,
    create_user uuid,
    is_delete boolean,
    delete_date timestamp with time zone,
    delete_user uuid,
    update_date timestamp with time zone,
    update_user uuid,
    brandid uuid,
    code character varying(200),
    title character varying(200),
    "user" character varying(200),
    password character varying(200),
    iorder integer,
    cooltype character varying(20)
);
    DROP TABLE public.t_model;
       public         heap    postgres    false            N           0    0    TABLE t_model    COMMENT     3   COMMENT ON TABLE public.t_model IS '矿机型号';
          public          postgres    false    215            O           0    0    COLUMN t_model.brandid    COMMENT     8   COMMENT ON COLUMN public.t_model.brandid IS '厂家id';
          public          postgres    false    215            P           0    0    COLUMN t_model.code    COMMENT     9   COMMENT ON COLUMN public.t_model.code IS '型号编号';
          public          postgres    false    215            Q           0    0    COLUMN t_model.title    COMMENT     :   COMMENT ON COLUMN public.t_model.title IS '型号名称';
          public          postgres    false    215            R           0    0    COLUMN t_model."user"    COMMENT     D   COMMENT ON COLUMN public.t_model."user" IS '默认矿机用户名';
          public          postgres    false    215            S           0    0    COLUMN t_model.password    COMMENT     C   COMMENT ON COLUMN public.t_model.password IS '默认矿机密码';
          public          postgres    false    215            T           0    0    COLUMN t_model.iorder    COMMENT     5   COMMENT ON COLUMN public.t_model.iorder IS '排序';
          public          postgres    false    215            U           0    0    COLUMN t_model.cooltype    COMMENT     X   COMMENT ON COLUMN public.t_model.cooltype IS '冷却类型，风冷，水冷，油冷';
          public          postgres    false    215            �            1259    24584    t_user    TABLE     V  CREATE TABLE public.t_user (
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
       public         heap    postgres    false            V           0    0    COLUMN t_user.user_name    COMMENT     :   COMMENT ON COLUMN public.t_user.user_name IS '用户名';
          public          postgres    false    216            W           0    0    COLUMN t_user.password    COMMENT     6   COMMENT ON COLUMN public.t_user.password IS '密码';
          public          postgres    false    216                      0    24598    t_boxes 
   TABLE DATA           �   COPY public.t_boxes (id, create_date, create_user, is_delete, delete_date, delete_user, update_date, update_user, lineid, total_miners, offline_miners, normal_miners, empty_miners, total_place, ip_start, ip_end, islift, empty_place) FROM stdin;
    public          postgres    false    218   �\                 0    24603    t_brand 
   TABLE DATA           �   COPY public.t_brand (id, create_date, create_user, is_delete, delete_date, delete_user, update_date, update_user, code, title, iorder, password, username) FROM stdin;
    public          postgres    false    219   y]                 0    24591    t_line 
   TABLE DATA           �  COPY public.t_line (id, create_date, create_user, is_delete, delete_date, delete_user, update_date, update_user, code, title, boxes, ip, rule_floors, rule_miners, rule_box_columns, actual_total_miners, actual_empty_place, actual_empty_miners, actual_normal_miners, actual_offline_miners, islift, standfloor, ip_start, miner_brandid, rule_username, rule_password, box_begin) FROM stdin;
    public          postgres    false    217   F^                 0    24578    t_miner 
   TABLE DATA           �   COPY public.t_miner (id, create_date, create_user, is_delete, delete_date, delete_user, update_date, update_user, status, brand, username, password, boxid, default_username, default_password, ip, cols, rows, child_box) FROM stdin;
    public          postgres    false    214   �^                 0    24581    t_model 
   TABLE DATA           �   COPY public.t_model (id, create_date, create_user, is_delete, delete_date, delete_user, update_date, update_user, brandid, code, title, "user", password, iorder, cooltype) FROM stdin;
    public          postgres    false    215   1t                 0    24584    t_user 
   TABLE DATA           �   COPY public.t_user (id, create_date, create_user, is_delete, delete_date, delete_user, update_date, update_user, user_name, password) FROM stdin;
    public          postgres    false    216   �t       �           2606    24602    t_boxes t_boxes_pkey 
   CONSTRAINT     R   ALTER TABLE ONLY public.t_boxes
    ADD CONSTRAINT t_boxes_pkey PRIMARY KEY (id);
 >   ALTER TABLE ONLY public.t_boxes DROP CONSTRAINT t_boxes_pkey;
       public            postgres    false    218            �           2606    24607    t_brand t_brand_pkey 
   CONSTRAINT     R   ALTER TABLE ONLY public.t_brand
    ADD CONSTRAINT t_brand_pkey PRIMARY KEY (id);
 >   ALTER TABLE ONLY public.t_brand DROP CONSTRAINT t_brand_pkey;
       public            postgres    false    219            �           2606    24597    t_line t_line_pkey 
   CONSTRAINT     P   ALTER TABLE ONLY public.t_line
    ADD CONSTRAINT t_line_pkey PRIMARY KEY (id);
 <   ALTER TABLE ONLY public.t_line DROP CONSTRAINT t_line_pkey;
       public            postgres    false    217            z           2606    24588    t_miner t_miner_pkey 
   CONSTRAINT     R   ALTER TABLE ONLY public.t_miner
    ADD CONSTRAINT t_miner_pkey PRIMARY KEY (id);
 >   ALTER TABLE ONLY public.t_miner DROP CONSTRAINT t_miner_pkey;
       public            postgres    false    214            |           2606    24611    t_model t_model_pkey 
   CONSTRAINT     R   ALTER TABLE ONLY public.t_model
    ADD CONSTRAINT t_model_pkey PRIMARY KEY (id);
 >   ALTER TABLE ONLY public.t_model DROP CONSTRAINT t_model_pkey;
       public            postgres    false    215            ~           2606    24590    t_user t_user_pkey 
   CONSTRAINT     P   ALTER TABLE ONLY public.t_user
    ADD CONSTRAINT t_user_pkey PRIMARY KEY (id);
 <   ALTER TABLE ONLY public.t_user DROP CONSTRAINT t_user_pkey;
       public            postgres    false    216               �   x�ݐA
1E��)����Y���&i�	z�t7�Y�Ax�c��=�����6�օwK�# "%`�H+ZM�Ԃ�e�?��Cp2~
��%Wc��Q2�+��.�e���,�zvx-!�F�.�iΐgqhc�&�h�#^���z�b��
yq         �   x�u�=
�@�:9�ٿ$����&��
A��^�F�h� x���$�0������O��V��	�'��P;��M<!,�gn��պM���DݹnN��:F��TRc�g�3���3/=���j .ʠ�ա��ͭ�\���(�▁"��Ȩ���LXC#�����,��n����
4t>�����U�         �   x���;1D��)���d�P!!N��D���t�p�A���b����֫٢d�f%62���UݴX����$P�Ara+a$֢\b�����~��u4�H���v��^��^=	���<0!�� q�wν�=A3            x��M�%����~WQ��DQ����zB�h��(���佰�|���+���03X~�D�#���F�����(H+
�El��+��!#��o��J���WM��o�����_�͏�����<�_��������������������Ғ��BXm6��*�@Kj�t���/��������KWټ�A��H���>[)f"�]z��_��GcZ�ӠΖ���K3'�A:7�����镯�7�ni�2���y�fh�eK���_���?�}ћ^AmD� 1u3b(�z�\��{��U�����X����:p9�)E{��_����赯��ǔ�$!AE�SA�&Xy[�m��r���^wv7^X����[�H��@)�s��4n�8���;����jã�2��`6�l/K�����ܠ�
�C��(�
<���yh"F�w��&[Nt���i��0��zI(Oc&u�q���G���R��8}�u��ۦk��!���u����n�f*F	p5�;pZ��y����Bm��DZih�}��P��qW<׮���w⋕/����L�A[R�j"}m�v��+S��/�_�z���bD�{�����2h�>gOSZY�n�8�5����KY�aҲ�w���͠E&u��n�w��_��O9��0� O�y����p�����;�G��׳�Zz��iH��) `��|7)w����#/Eڌs��Ȩ1ye�݁=iI5u5�߱˜<��{�tZ��*i9�������Y4��\�]|���N�y�m�|�2y��c���w�et�_����]c�/���/��(��Vw�-�o�8����/�Y7j�
-%��t��ɬ�Rd�m����W{�[�l�"���h��M�j��x��+�N|񕣆h�k���}�3�*`X�NZM�_|��o	|�Z&s����h�߇ÀהŻ��]|��U�����6z�R����rW�h��_�k�/�'>��/�I�¢C���"�yA[�&�F,w���������a�&P9־]#mN6S�<���m>�3�G�w��T��*Ђ�O^�Q��"EQ�������H���ŗq��A}����rʶn�|���VዶU�i�(����Kƭ{���w��>�U��KM7�:�'.}�tJP�[A��[��=�}���7N�X1�����C��@M���.Jt��>�U>y�������q�}z"�TM̊��*�ߧ�
_��">�x�������Ȍ�j�a��O}�lN��/fˎ��.�Q!Tk�?�t#��S`�/�4��(���wh�	�?s�zCǉ�Sa�/O�GN� �A@Sɂ���y�Ӿ�j�TX�7�6Kht�����ѧê��w�������ț����_a�w\�DHI�̜z-���.�ߧ���>��h�@Ǩ::�h�>�wOI{��~��ߧ�
_��;.�����G3���2��5�~�8�}*�<�K���6P���
�uO�k��|Uǉ�Sa���qwh�0\g �`�MG�JI=�^|�O��J�k5�BGN��a��{�uCǉ�Sa��f7^����C�b�F�s��Q�'�O��f�Z��:���P�eƎ�ܳS�rŜ�>V�rD���x]���I��>q�VS�z׾�4��
_���h���\u�.�jꡣ���$����Sa�/��7�2)NO���ކ侬�y�ߧ���m�T��T� ��@�j�]��`�<n�8�}*�<t8��~RZ�6w���}�n�fs�q��>V>��CX	�� ��g�%pN2����ǝ�'�O�����%E1s��#���'/�Q�N�ߧ�
_Y�hY0�����!0]�I�D�V�~��Sa���,I�C���S� �s��V6�ٲ���'�O��j��յ�6���\��2�#Ml3W-w���>V��U]{����o+~oߓ��y]�qv��TX�h�4
ۆ%������Z��<�ӻa��O�UyEi�ڠ奱a=\��1m{J�-��ˉ�SaU^�!ujt���/"o�˧m�#a�[�{��TX��-��~�Q��`��p�%�K���{��TX���
�5�9|�� 'I/����}'�O�Uy��RcD�ⷑE���/�K!۞����ߧª����cN�����6���M�T�6﹎o�}*�ʫy�2P;�]�������t�jH�~�=�}*�b�s��,��mv>�Dv�N��n��>V�U2b�<��{x�H�V��څ(����>���򼯣U�h����O}�>���[a��O�UyU3_��ݜ!-7�ї���D���i��SaU^s4U�
#E+��w�SN������ߧª�b�Z��A��--5��G<`t��.�k߉�SaU^�J��ϥ��4]�*��+F�p�X�7�>V�E�Ĉ=�J*@�������3[ۈ�o�8�}*�\uD�F�Le�8�*�Oׄݳ־�-�ߧª�2��[�(g 2(��N��?�6��ߧ���澫
!X'��u{�M!��5Nl�L8�}*�\��Ԉ���8����J3��{�e]3���M��O��'.j>y�?>�m����a��{��jxw\N|�
��JR�g�=B����}�+yG�W.�}�y߉�SaU^fm�ոsŎKe�i�կߖg���������ѷʪ$rM�}��^�P�)3]w��>V��;~�5f�!H�B��
���F�ʷ������*/������x���A��ШF;��k��{��TXy���֎�fO�W� ��^k9Bݣ��ߧª�$O��y [��3mFy@��0Ӱv׾o�}*��K}��1���Y_Cd��Y�U�-�=�}*�<��6b[ �� �P���c#�K4�;.��ɧª�F.ضXH�мe������jo	{����jmEG�V(Z��#�r�y��|�����TӀ�M=rpY��5����֯ث��?�u��,0�����z�29!'鿢	���=��Q�>� ⒃� ��kC4�_Q�3z��x��|i���U`��м7���+��~F�1D�]�&Ϙ��K��`��RKn�w|����p��G�q�{��r�����8O�3z�9`�s��nI���^ߎ�!��*�w����}ek�k�����W<X V��=n�w�{�w_�K�t�A�¡c� ,I���;_zg����ZwJ3�J����VÎ��2+�����s�T�	2���ͦ��>��Mk�_�K�C|�{ٸz�����Kg�z�̔랿b����s��a�[`�����F#m��\������{������n��/F������x��;����1$�Z����k�`��b̓J���U�'����Je��]R�����]�)�uI�w��s�:xi���kx�(�`l�6;��9~��1DU��y�2C�ַ;ʓ����q����s��S���j	O�p������Nr�ov�s��f�H���#o��G��.�ZO��_���s��
�M�<9���.a�5V��W��c�y�B��,:��l�Ɔ٢i�֛���s@Y�Z��D����D��\��//��cX)�̃!��q�K�cGa"�m5ջ���s�4�R����A.(.��:�Y<���*��s@�Cto���>��'/Mh8k�R���U�{��Y"]�-L���]a������1�f��ƪѲ�|���NJ�1���U?����1|�uW��k�$h4(���<f�����x��cl�A�J~�X��T�(�����!���[�<`��#�0)A��*Oahͫ:N|�9����������	�4A��e�9ݼ����2JY�r��#-�hCi	�u�[�q�{�g�9��P~[��0$&z�Ǜ��M\N|�9 f�ѩ���м�ՠ��Zne��Q�C|�9��\ey��;�q�UE��kl�=�������s�Y=��#V~[e�!=E��Xܱ�|#��1�~���͇ ���K��3�]S����������\+z%*�l-�N�����Y���ZE��
]%<vt �)�=l���=�=�>�Lq������4�K=\X��Z)��;�=怅��b�<�6W���@T�d�*���o�=急˖��y� >  蛍@�t�ym\b)��ep�{��%�hPߜb�tLМ'��q���лau�{�{�M�OV�0������w��s@*�[&�5.�;�䑤�Rm�t���s�6����Ǌ��T��(!Z5����{�{�Kʙ,��kX�7���Y�Q�ԫK������N���
s�G^^�o�H��,{c���@㇧s����adr�1c�e�`D[����{�����l�Kh�UR�tL����u�������
�A��8w���w�k�����P}�����s@^#�E}߻�HGw�p�8�'�[�����3�=Zx�D}߲VY�Z]q���=�v��cا��ї�0����A&������7�s�U��J�Z�당�U���m�%�2�U'���Q�+5�9�۵99Yeǟ:����W4��!��P=?}�A8o�^�2C�(� ۼ-H����Ήs��>KdY�����I�fl�<�.�9�O]�.���q�#��D�ȴ��g�o�=总h��'/O
_Y_�,����ؒ<�.7�;�=�<f�;E��K5Za����׾�y�rk\N|�9���5%��[�Y���6�ň�V'��p���i�
4%�C��+�H�7r�M\N|�9��;�����7��O*�\t�Y3�*��cG�G��ay7�+��nH���h#�V��{�W�q���V�Dg�H��g��z?���1$*��T�#��k^�3Z��*e{ ��|�9����躴;�8�A�4(#3�1��Yv��c��K��j8�G}�J�ɛi��K���c80Qoc J@���D(Y�6h�[��{�36.��0�@�G��8_��	5LS��=�=�q(&��>�B�z��-�N�<y��ռ'�?��IsO�Q���h�3�\��R�����|�9���=�J1e�\C��֬�!�����{��vӒԹ=t���r��Fu��˺i��9��=%�.K|i�(������8���N|�9���\�j�C��1�5U*:7ݵ����� k3�>w�@jS�(���-�L�wس��k�c8��*�Tf*amL�?�̙UsK;��Tt�{�1��0�q�Cu�ج�8�;dvs�{E�7�s�*���w��{�z����s���cl-�~�=�=�(y�.���9N�h�Q�D���w�;�=���U�/���oSϘ��\<������=�=怞��i�|ŋ*�ix��n���t�C��s�fU�w�-�׈�*�t�
�����{�{�y���`;J�[h^�
+�"��o�|�{��X����E��D\d٣W�m뎾�cHœ��)_�h�S[���ܴ6ID�n'�9��������!�H         s   x�K��0LNIL�5L�L�5IM2�M�H����,���R,9c���T3"tBpQ~~	� k�"��Ґ2k� �M��_���|8�n}�1��;fV$c�1�`tkc���� �1��         ]   x�K2J�4INK�57J5�52JM�M�43�5021JL642000��Â��9�-,��PL04�b�!v�J�8ML͈1�#2J9���b���� 6�1)     