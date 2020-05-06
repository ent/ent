FROM mcr.microsoft.com/mssql/server:2017-CU8-ubuntu

# Needed to be set ealier for script
ENV SA_PASSWORD yourStrong(!)Password

WORKDIR /usr/src/app/
ADD entrypoint.sh import-data.sh setup.sql /usr/src/app/

# Grant permissions for the import-data script to be executable
RUN chmod +x /usr/src/app/import-data.sh

CMD /bin/bash ./entrypoint.sh