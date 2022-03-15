const {createProxyMiddleware} = require("http-proxy-middleware");
module.exports = app => {
    app.use(
        createProxyMiddleware("/registration",{
            target: "http://localhost:8888",
            changeOrigin:  true
        })
    )
}