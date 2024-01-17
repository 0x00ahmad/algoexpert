import shutil
import os


dirs = os.listdir("./")
# each directory has a name that starts with x. <dirname>
# we want to remove the x. from the name via renaming

for each in dirs:
    if os.path.isdir(each) and ". " in each:
        new_name = each[2:]
        shutil.move(each, new_name)
        print(f"Renamed {each} to {new_name}")


