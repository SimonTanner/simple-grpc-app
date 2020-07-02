FROM migrate/migrate

ADD ./db/wait-for.sh .

ADD ./db/migrations ./migrations

# ADD ./db/docker-entrypoint.sh .

RUN chmod +x ./wait-for.sh

# RUN chmod +x ./docker-entrypoint.sh