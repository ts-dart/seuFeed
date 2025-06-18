import { NextConfig } from "next";

const nextConfig: NextConfig = {
  images: {
    remotePatterns: [
      {
        protocol: undefined, // aceita http e https
        hostname: '**',      // aceita qualquer domínio
      },
    ],
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
