/** @type {import('next').NextConfig} */
const nextConfig = {
  experimental: {
    serverActions: true,
  },
  async rewrites() {
    return [
      {
        source: "/api/:path*",
        destination: `${
          process.env.API_URL || process.env.NEXT_PUBLIC_API_URL
        }/v1/:path*`,
      },
    ];
  },
};

module.exports = nextConfig;
