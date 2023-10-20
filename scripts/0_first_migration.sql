CREATE TABLE IF NOT EXISTS delivery_info (
                                             id Serial primary key,
                                             name VARCHAR(50),
                                             phone VARCHAR(20),
                                             zip VARCHAR(10),
                                             city  VARCHAR(20),
                                             address VARCHAR(50),
                                             region VARCHAR(20),
                                             email VARCHAR(20)
);

CREATE TABLE IF NOT EXISTS payment_info (
                                            id Serial primary key,
                                            transaction VARCHAR(50),
                                            request_id VARCHAR(50),
                                            currency VARCHAR(8),
                                            amount Float,
                                            payment_dt Integer,
                                            bank VARCHAR(20),
                                            delivery_cost Float,
                                            goods_total Integer,
                                            custom_fee Float,
                                            provider VARCHAR(20)
);

CREATE TABLE IF NOT EXISTS item (
                                    chrt_id Integer primary key,
                                    track_number VARCHAR(30),
                                    price Integer,
                                    rid  VARCHAR(30),
                                    name VARCHAR(30),
                                    sale Integer,
                                    size VARCHAR(20),
                                    total_price Integer,
                                    nm_id Integer,
                                    brand VARCHAR(30),
                                    status Integer
);

CREATE TABLE IF NOT EXISTS orders (
                                      order_uid VARCHAR(50) primary key,
                                      entry VARCHAR(10),
                                      delivery Integer,
                                      payment Integer,
                                      CONSTRAINT fk_delivery
                                          FOREIGN KEY(delivery)
                                              REFERENCES delivery_info(id),
                                      CONSTRAINT fk_payment
                                          FOREIGN KEY(payment)
                                              REFERENCES payment_info(id),
                                      locale VARCHAR(10),
                                      internal_signature VARCHAR(20),
                                      customer_id VARCHAR(50),
                                      delivery_service VARCHAR(20),
                                      shardkey VARCHAR(10),
                                      sm_id Integer,
                                      date_created Timestamp,
                                      oof_shard VARCHAR(10),
                                      track_number VARCHAR(30)
);

CREATE TABLE IF NOT EXISTS item_order (
                                          id Serial primary key,
                                          chrt_id    Integer REFERENCES item (chrt_id) ON UPDATE CASCADE ON DELETE CASCADE,
                                          order_uid VARCHAR(20) REFERENCES orders (order_uid) ON UPDATE CASCADE
);
