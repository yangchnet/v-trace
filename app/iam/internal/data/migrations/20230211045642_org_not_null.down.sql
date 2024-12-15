ALTER TABLE `org` MODIFY `org_name` varchar(255);
ALTER TABLE `org` MODIFY `org_code` varchar(255) UNIQUE;
ALTER TABLE `org` MODIFY `legal_person_name` varchar(255);
ALTER TABLE `org` MODIFY `legal_person_phone` varchar(255);