output "Public ip" {
  value = "${digitalocean_droplet.api.ipv4_address}"
}

output "Name" {
  value = "${digitalocean_droplet.api.name}"
}

output "Hour Price" {
  value = "${digitalocean_droplet.api.price_hourly}"
}

output "Montly Price" {
  value = "${digitalocean_droplet.api.price_monthly}"
}
