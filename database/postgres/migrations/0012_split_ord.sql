-- +goose Up
CREATE TABLE payments
(
    id               varchar(255) NOT NULL PRIMARY KEY,
    order_id         varchar(255) NOT NULL,
    card_id          varchar(255) NOT NULL,
    provider         varchar(255) NOT NULL,
    transaction_id   text,
    idempotency_key  varchar(255),
    type             varchar(255),
    status           varchar(255) NOT NULL,
    fee_amount       numeric,
    fee_currency     varchar(100),
    properties       jsonb,
    error            jsonb,
    avs              text,
    created_at       timestamp with time zone DEFAULT NOW() NOT NULL,
    updated_at       timestamp with time zone
);

CREATE INDEX payments_order_id_idx
    on payments (order_id);

CREATE UNIQUE INDEX payments_transaction_id_idx
    on payments (transaction_id);

CREATE UNIQUE INDEX payments_idempotency_key_idx
    on payments (idempotency_key);

INSERT INTO payments(
    id,
    order_id,
    card_id,
    provider,
    idempotency_key,
    type,
    status,
    transaction_id,
    fee_amount,
    fee_currency,
    error,
    avs,
    created_at,
  syntax_error_for_generating_problem,
    updated_at
) SELECT
         id,
         id,
         card_id,
         'circle',
         idempotency_key,
         'non_3ds',
         checkout_status,
         checkout_id,
         payment_fee_amount,
         payment_fee_currency,
         payment_err_params,
         payment_avs_check_result,
         checkout_start,
         checkout_end
FROM orders WHERE failover_parent_order_id IS NULL;

UPDATE payments p SET order_id = o.id FROM orders o WHERE o.failover_parent_order_id = p.order_id;

-- +goose Down
DROP INDEX IF EXISTS payments_order_id_idx;

DROP INDEX IF EXISTS payments_transaction_id_idx;

DROP INDEX IF EXISTS payments_idempotency_key_idx;

DROP TABLE IF EXISTS payments;
