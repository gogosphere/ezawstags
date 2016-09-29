# ezawstags

```
taggerator account_role instanceid key:name key:name [...]
```

You can specify however many key:name pairs as you want.  If you use an existing key it will overwrite the name. The only requirements are the account_role, instanceid, and at least one key:name pair.

account_role = name of the account_role in your ~/.aws/credentials

instanceid = the actual instanceid of the resource e.g. i-6js79f75

key:name e.g.  bob:uncle ted:"koppel is a waffle" blah:bleh
