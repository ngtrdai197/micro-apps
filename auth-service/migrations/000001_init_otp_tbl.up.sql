CREATE TABLE otp
(
    id           SERIAL PRIMARY KEY,
    user_id      UUID        NOT NULL,
    phone_number TEXT        NOT NULL,
    otp_code     VARCHAR(6)  NOT NULL,
    is_valid     BOOLEAN     DEFAULT TRUE,
    expires_at   TIMESTAMPTZ NOT NULL,
    created_at   TIMESTAMPTZ DEFAULT now(),
    updated_at   TIMESTAMPTZ DEFAULT now(),
    deleted_at   TIMESTAMPTZ
);

CREATE INDEX idx_otp_phone_number ON public."otp" (phone_number);