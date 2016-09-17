Database related functions are stored here

##Database Layout

Table Name | Data Type | Explanation
------------ | ------------- | -------------
linkid | CHAR(6) | Unique key (path value in the url)
created | DATE | Date the link was added to the Database
userid | CHAR(10) | A semi persistant key which references a user who created the link
abuse_flags | SMALLINT | Number of times this link has been reported as abusive
clicks | INT | Times this link has been clicked
expires | DATE | An optional value stating when this link should expires
