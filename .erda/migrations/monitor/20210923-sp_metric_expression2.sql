UPDATE `sp_metric_expression` SET `enable` = 0 WHERE `id` in (SELECT `id` FROM (SELECT `id` FROM `sp_metric_expression` WHERE `enable` = 1 AND `expression` LIKE "%\"alias\":\"service_node\"%") AS `ids`);
INSERT `sp_metric_expression`(`attributes`,`expression`,`version`) VALUES("{}","{\"alias\":\"service_node\",\"filter\":{},\"functions\":[{\"aggregator\":\"sum\",\"field\":\"count\",\"field_script\":\"function invoke(field, tag){ return 1; }\"}],\"group\":[\"terminus_key\",\"msp_env_id\",\"service_id\"],\"metric\":\"application_service_node\",\"outputs\":[\"metric\"],\"select\":{\"application_id\":\"#application_id\",\"application_name\":\"#application_name\",\"instrumentation_library\":\"#instrumentation_library\",\"instrumentation_library_version\":\"#instrumentation_library_version\",\"msp_env_id\":\"#msp_env_id\",\"project_id\":\"#project_id\",\"project_name\":\"#project_name\",\"runtime_id\":\"#runtime_id\",\"runtime_name\":\"#runtime_name\",\"service_id\":\"#service_id\",\"service_name\":\"#service_name\",\"terminus_key\":\"#terminus_key\",\"workspace\":\"#workspace\"},\"window\":2}","3.0");