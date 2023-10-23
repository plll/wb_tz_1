CREATE TABLE IF NOT EXISTS delivery_info (
                                             id Serial primary key,
                                             name VARCHAR(50) NOT NULL,
                                             phone VARCHAR(20) NOT NULL,
                                             zip VARCHAR(10) NOT NULL,
                                             city  VARCHAR(20) NOT NULL,
                                             address VARCHAR(50) NOT NULL,
                                             region VARCHAR(20) NOT NULL,
                                             email VARCHAR(20) NOT NULL
);

CREATE TABLE IF NOT EXISTS payment_info (
                                            id Serial primary key,
                                            transaction VARCHAR(50) NOT NULL,
                                            request_id VARCHAR(50) NOT NULL,
                                            currency VARCHAR(8) NOT NULL,
                                            amount Float NOT NULL,
                                            payment_dt Integer NOT NULL,
                                            bank VARCHAR(20) NOT NULL,
                                            delivery_cost Float NOT NULL,
                                            goods_total Integer NOT NULL,
                                            custom_fee Float NOT NULL,
                                            provider VARCHAR(20) NOT NULL
);

CREATE TABLE IF NOT EXISTS item (
                                    chrt_id Integer primary key,
                                    track_number VARCHAR(30) NOT NULL,
                                    price Integer NOT NULL,
                                    rid  VARCHAR(30) NOT NULL,
                                    name VARCHAR(30) NOT NULL,
                                    sale Integer NOT NULL,
                                    size VARCHAR(20) NOT NULL,
                                    total_price Integer NOT NULL,
                                    nm_id Integer NOT NULL,
                                    brand VARCHAR(30) NOT NULL,
                                    status Integer NOT NULL
);

CREATE TABLE IF NOT EXISTS orders (
                                      order_uid VARCHAR(50) primary key,
                                      entry VARCHAR(10) NOT NULL,
                                      delivery Integer NOT NULL,
                                      payment Integer NOT NULL,
                                      CONSTRAINT fk_delivery
                                          FOREIGN KEY(delivery)
                                              REFERENCES delivery_info(id),
                                      CONSTRAINT fk_payment
                                          FOREIGN KEY(payment)
                                              REFERENCES payment_info(id),
                                      locale VARCHAR(10) NOT NULL,
                                      internal_signature VARCHAR(20) NOT NULL,
                                      customer_id VARCHAR(50) NOT NULL,
                                      delivery_service VARCHAR(20) NOT NULL,
                                      shardkey VARCHAR(10) NOT NULL,
                                      sm_id Integer NOT NULL,
                                      date_created Timestamp NOT NULL,
                                      oof_shard VARCHAR(10) NOT NULL,
                                      track_number VARCHAR(30) NOT NULL
);

CREATE TABLE IF NOT EXISTS item_order (
                                          id Serial primary key,
                                          chrt_id    Integer REFERENCES item (chrt_id) ON UPDATE CASCADE ON DELETE CASCADE,
                                          order_uid VARCHAR(20) REFERENCES orders (order_uid) ON UPDATE CASCADE
);
