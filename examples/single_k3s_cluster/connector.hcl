container "local_connector" {
  image {
    name = "shipyardrun/connector:v0.0.1"
  }

  env_var = {
    "BIND_ADDR_GRPC" = "0.0.0.0:9090"
    "BIND_ADDR_HTTP" = "0.0.0.0:9091"
    "LOG_LEVEL" = "debug"
  }
  
  port_range {
    range = "9090-9091"
    enable_host = true
  }

  port_range {
    range = "12000-12010"
    enable_host = true
  }

  network {
    name = "network.cloud"
  }
}
