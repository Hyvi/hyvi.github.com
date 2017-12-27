import fabric
from fabric.api import env, run
env.hosts = ['120.24.242.232']
env.port = 22
env.user = 'root'
def blog_gliese():
    fabric.contrib.project.rsync_project(remote_dir='/mnt/blog/', local_dir='./_site/', ssh_opts='')
