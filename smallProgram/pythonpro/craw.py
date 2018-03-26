import os

dir=r'/var/lib/docker/workfile/testoverlay/lower'

print dir

for root, dirs, files in os.walk(dir):
    #print "root",root,"dir",dirs,"files",files
    for f in files:
        try:
            fpath=str(os.path.join(root,f))
            print "fpath",fpath,"list",fpath.split(dir)[1]
        except UnicodeEncodeError as err:
            print err
            pass