-- Enable row-level security on the users table.
ALTER TABLE "users" ENABLE ROW LEVEL SECURITY;

-- Create a policy that restricts access to rows in the users table based on the current tenant.
CREATE POLICY tenant_isolation ON "users"
    USING ("tenant_id" = current_setting('app.current_tenant')::integer);
