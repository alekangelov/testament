import React from "react";

interface P extends React.SVGProps<SVGSVGElement> {
  size?: number | string;
}

export function Apple({ size = 24, ...props }: P) {
  return (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      width={size}
      height={size}
      fill="none"
      viewBox="0 0 128 157"
      {...props}
    >
      <path
        fill="currentColor"
        d="M123.912 53.474C123 54.18 106.9 63.23 106.9 83.356c0 23.278 20.487 31.513 21.1 31.717-.094.502-3.255 11.278-10.802 22.259-6.729 9.662-13.757 19.309-24.449 19.309-10.691 0-13.443-6.196-25.785-6.196-12.028 0-16.305 6.4-26.084 6.4-9.78 0-16.604-8.941-24.45-19.921C7.343 124.03 0 103.999 0 84.987c0-30.494 19.874-46.666 39.433-46.666 10.393 0 19.056 6.808 25.581 6.808 6.21 0 15.896-7.216 27.72-7.216 4.48 0 20.581.408 31.178 15.56zm-36.791-28.47c4.89-5.789 8.348-13.82 8.348-21.851 0-1.114-.094-2.243-.298-3.153-7.956.298-17.421 5.286-23.129 11.89-4.48 5.082-8.663 13.114-8.663 21.255 0 1.223.204 2.447.299 2.839.503.094 1.32.204 2.138.204 7.138 0 16.116-4.769 21.305-11.184z"
      ></path>
    </svg>
  );
}

export function Facebook({ size = 24, ...props }: P) {
  return (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      x="0"
      y="0"
      enableBackground="new 0 0 40 40"
      version="1.1"
      viewBox="0 0 40 40"
      xmlSpace="preserve"
      width={size}
      height={size}
    >
      <linearGradient
        id="SVGID_1_"
        x1="-277.375"
        x2="-277.375"
        y1="406.602"
        y2="407.573"
        gradientTransform="matrix(40 0 0 -39.7778 11115.001 16212.334)"
        gradientUnits="userSpaceOnUse"
      >
        <stop offset="0" stopColor="#0062E0"></stop>
        <stop offset="1" stopColor="#19AFFF"></stop>
      </linearGradient>
      <path
        fill="url(#SVGID_1_)"
        d="M16.7 39.8C7.2 38.1 0 29.9 0 20 0 9 9 0 20 0s20 9 20 20c0 9.9-7.2 18.1-16.7 19.8l-1.1-.9h-4.4l-1.1.9z"
      ></path>
      <path
        fill="#fff"
        d="M27.8 25.6l.9-5.6h-5.3v-3.9c0-1.6.6-2.8 3-2.8H29V8.2c-1.4-.2-3-.4-4.4-.4-4.6 0-7.8 2.8-7.8 7.8V20h-5v5.6h5v14.1c1.1.2 2.2.3 3.3.3 1.1 0 2.2-.1 3.3-.3V25.6h4.4z"
      ></path>
    </svg>
  );
}

export function Google({ size = 24, ...props }: P) {
  return (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      width={size}
      height={size}
      viewBox="0 0 24 24"
    >
      <path
        fill="#4285F4"
        d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z"
      ></path>
      <path
        fill="#34A853"
        d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z"
      ></path>
      <path
        fill="#FBBC05"
        d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z"
      ></path>
      <path
        fill="#EA4335"
        d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z"
      ></path>
      <path fill="none" d="M1 1h22v22H1z"></path>
    </svg>
  );
}

export function Close({ size = 14, ...props }: P) {
  return (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      width={size}
      height={size}
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      stroke-width="2"
      stroke-linecap="round"
      stroke-linejoin="round"
      {...props}
    >
      <line x1="18" y1="6" x2="6" y2="18" />
      <line x1="6" y1="6" x2="18" y2="18" />
    </svg>
  );
}

export function Check({ size = 14, ...props }: P) {
  return (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      width={size}
      height={size}
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      stroke-width="2"
      stroke-linecap="round"
      stroke-linejoin="round"
      {...props}
    >
      <polyline points="20 6 9 17 4 12" />
    </svg>
  );
}

export function Info({ size = 14, ...props }: P) {
  return (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      width={size}
      height={size}
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      stroke-width="2"
      stroke-linecap="round"
      stroke-linejoin="round"
      {...props}
    >
      <circle cx="12" cy="12" r="10" />
      <line x1="12" y1="16" x2="12" y2="12" />
      <line x1="12" y1="8" x2="12.01" y2="8" />
    </svg>
  );
}

export function Warning({ size = 14, ...props }: P) {
  return (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      stroke-width="2"
      stroke-linecap="round"
      stroke-linejoin="round"
      width={size}
      height={size}
      {...props}
    >
      <path d="M10.29 3.86L1.82 18a2 2 0 0 0 1.71 3h16.94a2 2 0 0 0 1.71-3L13.71 3.86a2 2 0 0 0-3.42 0z" />
      <line x1="12" y1="9" x2="12" y2="13" />
      <line x1="12" y1="17" x2="12.01" y2="17" />
    </svg>
  );
}

export function Error({ size = 14, ...props }: P) {
  return (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      width={size}
      height={size}
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      stroke-width="2"
      stroke-linecap="round"
      stroke-linejoin="round"
      {...props}
    >
      <polygon points="7.86 2 16.14 2 22 7.86 22 16.14 16.14 22 7.86 22 2 16.14 2 7.86 7.86 2" />
      <line x1="15" y1="9" x2="9" y2="15" />
      <line x1="9" y1="9" x2="15" y2="15" />
    </svg>
  );
}
