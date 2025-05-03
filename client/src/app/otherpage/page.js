import Link from "next/link";

export default function OtherPage() {
  return (
    <div>
      <div className="mb-8">
        <h1 className="text-3xl font-bold text-slate-800 mb-2">About</h1>
        <p className="text-slate-500">Learn more about this application</p>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
        <div className="bg-white rounded-xl border border-slate-200 shadow-sm overflow-hidden">
          <div className="p-6">
            <h2 className="text-xl font-semibold text-slate-800 mb-4">Project Overview</h2>
            <p className="text-slate-600 mb-4">
              This application demonstrates a Fibonacci calculator implemented using Next.js and Docker. It shows how to build modern web applications with a microservices
              architecture.
            </p>
            <p className="text-slate-600 mb-6">The application uses React hooks, modern UI design principles, and efficient state management to provide a great user experience.</p>
            <Link href="/" className="inline-flex items-center px-4 py-2 bg-blue-50 text-blue-700 rounded-lg hover:bg-blue-100 transition-colors font-medium text-sm">
              <svg
                className="mr-2"
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
                <line x1="19" y1="12" x2="5" y2="12"></line>
                <polyline points="12 19 5 12 12 5"></polyline>
              </svg>
              Return to Calculator
            </Link>
          </div>
        </div>

        <div className="bg-white rounded-xl border border-slate-200 shadow-sm overflow-hidden">
          <div className="p-6">
            <h2 className="text-xl font-semibold text-slate-800 mb-4">Technologies Used</h2>
            <ul className="space-y-3">
              <li className="flex items-start">
                <div className="rounded-full p-1 bg-green-100 text-green-800 mr-3 mt-0.5">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    width="14"
                    height="14"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    strokeWidth="2"
                    strokeLinecap="round"
                    strokeLinejoin="round"
                  >
                    <polyline points="20 6 9 17 4 12"></polyline>
                  </svg>
                </div>
                <div>
                  <span className="font-medium text-slate-800">Next.js</span>
                  <p className="text-sm text-slate-500">React framework with server-side rendering</p>
                </div>
              </li>
              <li className="flex items-start">
                <div className="rounded-full p-1 bg-green-100 text-green-800 mr-3 mt-0.5">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    width="14"
                    height="14"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    strokeWidth="2"
                    strokeLinecap="round"
                    strokeLinejoin="round"
                  >
                    <polyline points="20 6 9 17 4 12"></polyline>
                  </svg>
                </div>
                <div>
                  <span className="font-medium text-slate-800">Docker</span>
                  <p className="text-sm text-slate-500">Containerization for consistent deployment</p>
                </div>
              </li>
              <li className="flex items-start">
                <div className="rounded-full p-1 bg-green-100 text-green-800 mr-3 mt-0.5">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    width="14"
                    height="14"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    strokeWidth="2"
                    strokeLinecap="round"
                    strokeLinejoin="round"
                  >
                    <polyline points="20 6 9 17 4 12"></polyline>
                  </svg>
                </div>
                <div>
                  <span className="font-medium text-slate-800">Tailwind CSS</span>
                  <p className="text-sm text-slate-500">Utility-first CSS framework</p>
                </div>
              </li>
            </ul>
          </div>
        </div>
      </div>
    </div>
  );
}
