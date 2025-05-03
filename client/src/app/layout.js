import { Inter } from "next/font/google";
import "./globals.css";
import Link from "next/link";

const inter = Inter({ subsets: ["latin"] });

export const metadata = {
  title: "Fibonacci Calculator",
  description: "A Next.js implementation of the Fibonacci calculator",
};

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      <body className={`${inter.className} bg-white`}>
        <div className="flex flex-col min-h-screen max-w-6xl mx-auto">
          <header className=" border-b border-slate-200 sticky top-0 z-10 backdrop-blur-sm bg-white/90">
            <div className="container mx-auto px-4 py-4 flex justify-between items-center">
              <div className="flex items-center gap-2">
                <div className="w-10 h-10 rounded-lg bg-gradient-to-br from-blue-500 to-indigo-600 flex items-center justify-center text-white">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    width="24"
                    height="24"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    strokeWidth="2"
                    strokeLinecap="round"
                    strokeLinejoin="round"
                  >
                    <path d="M6 9v12m6-16v16m6-8v8"></path>
                  </svg>
                </div>
                <span className="font-semibold text-xl text-slate-800">FibCalc</span>
              </div>

              <nav className="flex gap-4">
                <Link href="/" className="px-4 py-2 rounded-md hover:bg-slate-100 transition-colors text-slate-700 font-medium">
                  Home
                </Link>
                <Link href="/otherpage" className="px-4 py-2 rounded-md hover:bg-slate-100 transition-colors text-slate-700 font-medium">
                  About
                </Link>
              </nav>
            </div>
          </header>

          <main className="flex-grow p-6 md:p-8 container mx-auto">{children}</main>

          <footer className="bg-white border-t border-slate-200 py-6">
            <div className="container mx-auto px-4">
              <div className="flex flex-col md:flex-row justify-between items-center gap-4">
                <div className="flex items-center gap-2">
                  <div className="w-8 h-8 rounded-lg bg-gradient-to-br from-blue-500 to-indigo-600 flex items-center justify-center text-white">
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      width="16"
                      height="16"
                      viewBox="0 0 24 24"
                      fill="none"
                      stroke="currentColor"
                      strokeWidth="2"
                      strokeLinecap="round"
                      strokeLinejoin="round"
                    >
                      <path d="M6 9v12m6-16v16m6-8v8"></path>
                    </svg>
                  </div>
                  <span className="font-medium text-slate-700">FibCalc</span>
                </div>
                <p className="text-slate-500 text-sm">Built with Next.js • {new Date().getFullYear()}</p>
                <div className="flex gap-4">
                  <a href="https://nextjs.org" target="_blank" rel="noopener noreferrer" className="text-slate-500 hover:text-slate-800 transition-colors">
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      width="20"
                      height="20"
                      viewBox="0 0 24 24"
                      fill="none"
                      stroke="currentColor"
                      strokeWidth="2"
                      strokeLinecap="round"
                      strokeLinejoin="round"
                    >
                      <path d="M12 2L2 7l10 5 10-5-10-5zM2 17l10 5 10-5M2 12l10 5 10-5"></path>
                    </svg>
                  </a>
                  <a href="https://github.com" target="_blank" rel="noopener noreferrer" className="text-slate-500 hover:text-slate-800 transition-colors">
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      width="20"
                      height="20"
                      viewBox="0 0 24 24"
                      fill="none"
                      stroke="currentColor"
                      strokeWidth="2"
                      strokeLinecap="round"
                      strokeLinejoin="round"
                    >
                      <path d="M9 19c-5 1.5-5-2.5-7-3m14 6v-3.87a3.37 3.37 0 0 0-.94-2.61c3.14-.35 6.44-1.54 6.44-7A5.44 5.44 0 0 0 20 4.77 5.07 5.07 0 0 0 19.91 1S18.73.65 16 2.48a13.38 13.38 0 0 0-7 0C6.27.65 5.09 1 5.09 1A5.07 5.07 0 0 0 5 4.77a5.44 5.44 0 0 0-1.5 3.78c0 5.42 3.3 6.61 6.44 7A3.37 3.37 0 0 0 9 18.13V22"></path>
                    </svg>
                  </a>
                </div>
              </div>
            </div>
          </footer>
        </div>
      </body>
    </html>
  );
}
