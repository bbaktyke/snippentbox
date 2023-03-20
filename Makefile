run:
	docker run --rm --name snippetbox-mysql -e MYSQL_ROOT_PASSWORD=root -e MYSQL_USER=snippet -e MYSQL_PASSWORD=123 -e MYSQL_DATABASE=snippetbox -p 3307:3306 -d mysql
	sleep 20 && docker exec -i snippetbox-mysql mysql -uroot -proot snippetbox < mysql_dump.sql
exec:
	docker exec -it snippetbox-mysql mysql -uroot -proot
stop:
	docker exec -i snippetbox-mysql mysqldump -uroot -proot snippetbox > mysql_dump.sql
	docker stop snippetbox-mysql

# docker run -it --network some-network --rm mysql mysql -hsnippetbox-mysql -uexample-user -p
# docker exec -it some-mysql bash
# docker logs some-mysql

# MYSQL_ROOT_PASSWORD=123
# 	MYSQL_USER=snippet
# 	MYSQL_PASSWORD=123
# 	MYSQL_DATABASE=snippetbox
# sleep 20 && docker exec -i snippetbox-mysql mysql -uroot -proot snippetbox < mysql_dump.sql