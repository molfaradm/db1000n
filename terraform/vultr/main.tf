resource "vultr_ssh_key" "id_rsa" {
  name = "id_rsa"
  ssh_key = "${file("id_rsa.pub")}"
}

resource "vultr_instance" "db1000h" {

  count = 10

  plan = "vc2-1c-1gb"
  region = "sgp"
  os_id = "387"
  hostname = "db1000h-${count.index}"
  enable_ipv6 = false
  ddos_protection = false
  ssh_key_ids = ["${vultr_ssh_key.id_rsa.id}"]

  connection {
    host = self.main_ip
    user = "root"
    type = "ssh"
    private_key = file(var.pvt_key)
    timeout = "2m"
  }
  
  provisioner "remote-exec" {
    inline = [
      "export PATH=$PATH:/usr/bin",
      # install nginx
      "sudo apt update",
      "sudo apt install -y git wget",
      "wget https://github.com/Arriven/db1000n/releases/download/v0.7.12/db1000n-v0.7.12-linux-amd64.tar.gz -O /root/db1000n-v0.7.12-linux-amd64.tar.gz",
      "cd /root/; tar xzf db1000n-v0.7.12-linux-amd64.tar.gz",
      "printf '%s\n' '#!/bin/bash' '/root/db1000n &' 'exit 0' | sudo tee -a /etc/rc.local",
      "chmod +x /etc/rc.local",
      "systemctl start rc-local",
      "systemctl enable rc-local"
    ]
  }
}

