-- Function to audit changes in the users table.
CREATE OR REPLACE FUNCTION audit_users_changes()
RETURNS TRIGGER AS $$
BEGIN
    IF (TG_OP = 'INSERT') THEN
        INSERT INTO user_audit_logs(operation_type, operation_time, new_value)
        VALUES (TG_OP, CURRENT_TIMESTAMP, row_to_json(NEW));
        RETURN NEW;
    ELSIF (TG_OP = 'UPDATE') THEN
        INSERT INTO user_audit_logs(operation_type, operation_time, old_value, new_value)
        VALUES (TG_OP, CURRENT_TIMESTAMP, row_to_json(OLD), row_to_json(NEW));
        RETURN NEW;
    ELSIF (TG_OP = 'DELETE') THEN
        INSERT INTO user_audit_logs(operation_type, operation_time, old_value)
        VALUES (TG_OP, CURRENT_TIMESTAMP, row_to_json(OLD));
        RETURN OLD;
    END IF;
    RETURN NULL;
END;
$$ LANGUAGE plpgsql;

-- Trigger for INSERT operations.
CREATE TRIGGER users_insert_audit AFTER INSERT ON users FOR EACH ROW EXECUTE FUNCTION audit_users_changes();

-- Trigger for UPDATE operations.
CREATE TRIGGER users_update_audit AFTER UPDATE ON users FOR EACH ROW EXECUTE FUNCTION audit_users_changes();

-- Trigger for DELETE operations.
CREATE TRIGGER users_delete_audit AFTER DELETE ON users FOR EACH ROW EXECUTE FUNCTION audit_users_changes();
