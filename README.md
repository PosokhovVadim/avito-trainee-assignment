# avito-trainee-assignment
	docker exec -it postgres_banner_db createdb --username=postgres --owner=postgres banner_db
b569895d6430   avito_banner    "./main"                 2 minutes ago   Up 2 minutes               0.0.0.0:8090->8090/tcp, :::8090->8090/tcp   avito_banner
c9204b035394   postgres:16.0   "docker-entrypoint.s…"   2 minutes ago   Up 2 minutes (healthy)     0.0.0.0:5432->5432/tcp, :::5432->5432/tcp   postgres_banner_db
bb17d61f2434   base            "bash"                   2 minutes ago   Exited (0) 2 minutes ago                                               avito_dependencies_1