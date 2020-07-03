FROM migrate/migrate

ADD ./db/wait-for.sh .

ADD ./db/migrations ./migrations

RUN chmod +x ./wait-for.sh
