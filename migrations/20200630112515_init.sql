-- +goose Up
CREATE TABLE public.orders (
    id serial NOT NULL,
    created_at timestamp NOT NULL DEFAULT now(),
    CONSTRAINT orders_pk PRIMARY KEY (id)
);

CREATE TABLE public.goods (
    id serial NOT NULL,
    created_at timestamp NOT NULL DEFAULT now(),
    is_deleted bool NOT NULL DEFAULT false,
    name varchar(64) NOT NULL,
    CONSTRAINT goods_pk PRIMARY KEY (id)
);

CREATE TABLE public.order_items (
    id serial NOT NULL,
    created_at timestamp NOT NULL DEFAULT now(),
    is_deleted bool NOT NULL DEFAULT false,
    good_id int4 NOT NULL,
    quantity varchar NOT NULL,
    order_id int4 NOT NULL
);

ALTER TABLE public.order_items ADD CONSTRAINT order_items_good FOREIGN KEY (good_id) REFERENCES goods(id);
ALTER TABLE public.order_items ADD CONSTRAINT order_items_order FOREIGN KEY (order_id) REFERENCES orders(id);

-- +goose Down
DROP TABLE public.order_items;
DROP TABLE public.goods;
DROP TABLE public.orders;
