-- this is okay for the local testing purposes

-- total admin
CREATE USER 'toolset_admin'@'localhost' IDENTIFIED BY 'password';
GRANT all ON toolset.* TO 'toolset_admin'@'localhost';

-- only select
CREATE USER 'toolset_select'@'localhost' IDENTIFIED BY 'password';
GRANT select ON toolset.* TO 'toolset_select'@'localhost';

-- only insert
CREATE USER 'toolset_insert'@'localhost' IDENTIFIED BY 'password';
GRANT insert ON toolset.* TO 'toolset_insert'@'localhost';

-- only delete
CREATE USER 'toolset_delete'@'localhost' IDENTIFIED BY 'password';
GRANT delete ON toolset.* TO 'toolset_delete'@'localhost';

-- only update
CREATE USER 'toolset_update'@'localhost' IDENTIFIED BY 'password';
GRANT update ON toolset.* TO 'toolset_update'@'localhost';
