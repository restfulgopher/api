provider "digitalocean" {
  # You need to set this in your .bashrc
  # export DIGITALOCEAN_TOKEN="Your API TOKEN"
  #
}

resource "digitalocean_droplet" "api" {
  # Obtain your ssh_key id number via your account. See Document https://developers.digitalocean.com/documentation/v2/#list-all-keys
  ssh_keys           = [7559088]         #
  image              = "${var.centos}"
  region             = "${var.do_lon1}"
  size               = "512mb"
  private_networking = true
  backups            = true
  ipv6               = true
  name               = "api"

  provisioner "remote-exec" {
    inline = [
      "export PATH=$PATH:/usr/bin",
      "wget -qO- https://get.docker.com/ | sh",
      "yum install -y epel-release",
      "yum install -y python-pip",
      "pip install docker-compose",
      "yum upgrade -y python*",
      "systemctl start docker.service",
      "wget https://raw.githubusercontent.com/betalotest/api/master/docker-compose.yml",
      "docker-compose pull api",
      "docker-compose pull stub_api",
      "docker-compose pull spec",
      "docker-compose pull nginx",
      "docker-compose up -d nginx",
    ]

    connection {
      type     = "ssh"
      private_key = "${file("~/.ssh/id_rsa")}"
      user     = "root"
      timeout  = "2m"
    }
  }
}

resource "digitalocean_domain" "api" {
  name       = "api.alesr.me"
  ip_address = "${digitalocean_droplet.api.ipv4_address}"
}
