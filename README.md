# gologger

gologger - Simple webserver for testing purposes.

    When the server receives a GET request to path "/" displays the port on which the app is running.
       
    When the server receives a GET request to path "/logger" displays random transactional data.

Usage of gologger:
  
    -a string
        Hostname/IP address (default "0.0.0.0")  

    -p int
        Port (default 8080)

Example:

    gologger -a 192.168.99.78 -p 9090

    2019/08/15 17:22:05 Server started...

Transactional data fields are separated by pipes:

    datetime|id|user|operation|duration|status|code|codeDescription

Example:

    2019-08-07 10:10:51.223|1565190651223457650|ubbo-sathla|changeUser|548|SUCCESS|0|SUCCESS
    2019-08-07 10:10:51.385|1565190651385873762|nyarlathotep|cancelProduct|1214|FAILED|440|Not Supported