import Fib from "@/components/Fib";

export default function Home() {
  return (
    <div>
      <div className="mb-8">
        <h1 className="text-3xl font-bold text-slate-800 mb-2">Fibonacci Calculator</h1>
        <p className="text-slate-500">Calculate Fibonacci numbers quickly and efficiently</p>
      </div>
      <Fib />
    </div>
  );
}
