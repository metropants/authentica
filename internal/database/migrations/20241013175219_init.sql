-- +goose Up
-- +goose StatementBegin
create extension if not exists "uuid-ossp";

create or replace function set_updated_at()
    returns trigger as
$$
begin
    NEW.updated_at = now();
    return NEW;
end;
$$ language plpgsql;

create or replace function trigger_updated_at(table_name regclass)
    returns void as
$$
begin
    execute format('CREATE TRIGGER set_updated_at
        BEFORE UPDATE
        ON %s
        FOR EACH ROW
        WHEN (OLD is distinct from NEW)
    EXECUTE FUNCTION set_updated_at();', table_name);
end;
$$ language plpgsql;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop extension if exists "uuid-ossp";
-- +goose StatementEnd
