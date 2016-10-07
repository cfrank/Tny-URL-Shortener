Database related functions are stored here

##Database Layout

###Link Table
Table Name | Data Type | Explanation
------------ | ------------- | -------------
linkid | CHAR(6) | Unique key (path value in the url)
source | VARCHAR(2083) | The url which the user will be redirected to
created | BIGINT | Seconds since Epoch when the link was added to the Database
userid | CHAR(10) | A semi persistant key which references a user who created the link
abuseflags | SMALLINT | Number of times this link has been reported as abusive
clicks | INT | Times this link has been clicked
expires | BIGINT | An optional value stating when this link should expire (TODO)

###User Table
This table is needed because when someone vists a page they will automatically
be given a userid/key. If they decide they don't want to make a shortened url
then there will never be a record of their userid. But it will still be active in
their localstorage. So, if they decide to come back later and make a short url
their userid could in theory be given out to someone else...

Table Name | Data Type | Explanation
------------ | ------------- | -------------
userid | CHAR(10) | A semi persistant key which references a user who created the link