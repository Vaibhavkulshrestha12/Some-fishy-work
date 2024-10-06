
module.exports = function override(config, env) {
    if (env === 'development') {
      config.devServer.allowedHosts = ['localhost', 'your_ip_here'];
    }
    return config;
  };
  