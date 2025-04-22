import { NextConfig } from "next";

const nextConfig: NextConfig = {
    images: {
        domains: ['www.cnnbrasil.com.br', 's2-g1.glbimg.com', 'encrypted-tbn0.gstatic.com'],
    },
    redirects: async () => [
        {
            source: "/",
            destination: "/home",
            permanent: true,
        },
    ],
};

export default nextConfig;