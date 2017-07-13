
## status-page : Script to generate a html page with Instance status

### Note
Assumed you have setup the configuration file correctly and your instances tag `Name` is set
them this script should work.

An other important thing is that this was build for our neea,d so we based it on that a tag
has the following format of `([A-Za-z]|\-)*` follow `[0-9].*` ; Example prod-momo1, test-momo1.
If not then look in the `ec2status.go` file. Th reason is to be able to group the instances.


### Feedback
Feedback and bug report welcome...

Enjoy, Momo
