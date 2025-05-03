"use client";

import React, { useState, useEffect } from "react";
import axios from "axios";

const Fib = () => {
  const [seenIndexes, setSeenIndexes] = useState([]);
  const [values, setValues] = useState({});
  const [index, setIndex] = useState("");
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const [successMessage, setSuccessMessage] = useState(null);

  useEffect(() => {
    fetchValues();
    fetchIndexes();
  }, []);

  const fetchValues = async () => {
    try {
      const values = await axios.get("/api/values/current");
      setValues(values.data);
    } catch (err) {
      console.error("Error fetching values:", err);
    }
  };

  const fetchIndexes = async () => {
    try {
      const seenIndexes = await axios.get("/api/values/all");
      setSeenIndexes(seenIndexes.data);
    } catch (err) {
      console.error("Error fetching indexes:", err);
    }
  };

  const handleSubmit = async (event) => {
    event.preventDefault();

    if (!index.trim()) {
      setError("Please enter a number");
      return;
    }

    try {
      setLoading(true);
      setError(null);

      await axios.post("/api/values", {
        index: index,
      });

      setSuccessMessage("Calculation requested!");
      setTimeout(() => setSuccessMessage(null), 3000);

      setIndex("");
      // Refresh data after submission
      fetchValues();
      fetchIndexes();
    } catch (err) {
      setError("Error submitting calculation. Please try again.");
      console.error("Error submitting index:", err);
    } finally {
      setLoading(false);
    }
  };

  const renderSeenIndexes = () => {
    if (!seenIndexes || seenIndexes.length === 0) {
      return (
        <div className="flex items-center justify-center py-6 text-slate-500">
          <svg
            className="mr-2"
            xmlns="http://www.w3.org/2000/svg"
            width="18"
            height="18"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            strokeWidth="2"
            strokeLinecap="round"
            strokeLinejoin="round"
          >
            <path d="M22 12h-4l-3 9L9 3l-3 9H2" />
          </svg>
          No indices calculated yet
        </div>
      );
    }

    return (
      <div className="flex flex-wrap gap-2">
        {seenIndexes.map(({ number }) => (
          <span key={number} className="px-3 py-1 bg-blue-100 text-blue-700 rounded-full text-sm font-medium">
            {number}
          </span>
        ))}
      </div>
    );
  };

  const renderValues = () => {
    const entries = [];
    const keys = Object.keys(values);

    if (keys.length === 0) {
      return (
        <div className="bg-white border border-slate-200 rounded-lg p-6 text-center text-slate-500">
          <svg
            className="mx-auto mb-3"
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
            <circle cx="12" cy="12" r="10"></circle>
            <line x1="12" y1="8" x2="12" y2="12"></line>
            <line x1="12" y1="16" x2="12.01" y2="16"></line>
          </svg>
          No values calculated yet
        </div>
      );
    }

    for (let key in values) {
      entries.push(
        <div key={key} className="bg-white border border-slate-200 rounded-lg p-5 flex items-center justify-between mb-3 hover:shadow-md transition-shadow">
          <div>
            <span className="text-xs uppercase tracking-wider text-slate-500 font-semibold">Input</span>
            <div className="text-2xl font-bold text-blue-600">{key}</div>
          </div>
          <div className="bg-slate-100 w-12 h-12 flex items-center justify-center rounded-full">
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
              <polyline points="9 18 15 12 9 6"></polyline>
            </svg>
          </div>
          <div className="text-right">
            <span className="text-xs uppercase tracking-wider text-slate-500 font-semibold">Result</span>
            <div className="text-2xl font-bold text-indigo-600">{values[key]}</div>
          </div>
        </div>
      );
    }

    return entries;
  };

  return (
    <div className="grid grid-cols-1 lg:grid-cols-2 gap-8">
      <div>
        <div className="bg-white rounded-xl border border-slate-200 shadow-sm overflow-hidden">
          <div className="p-6 border-b border-slate-200">
            <h2 className="text-xl font-semibold text-slate-800">Calculate Fibonacci</h2>
            <p className="text-slate-500 mt-1">Enter an index to calculate its Fibonacci value</p>
          </div>

          <form onSubmit={handleSubmit} className="p-6">
            {error && (
              <div className="mb-4 p-3 bg-red-50 border border-red-200 text-red-700 rounded-lg text-sm flex items-center">
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
                  <circle cx="12" cy="12" r="10"></circle>
                  <line x1="12" y1="8" x2="12" y2="12"></line>
                  <line x1="12" y1="16" x2="12.01" y2="16"></line>
                </svg>
                {error}
              </div>
            )}

            {successMessage && (
              <div className="mb-4 p-3 bg-green-50 border border-green-200 text-green-700 rounded-lg text-sm flex items-center">
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
                  <path d="M22 11.08V12a10 10 0 1 1-5.93-9.14"></path>
                  <polyline points="22 4 12 14.01 9 11.01"></polyline>
                </svg>
                {successMessage}
              </div>
            )}

            <div className="mb-5">
              <label className="block mb-2 text-sm font-medium text-slate-700">Fibonacci Index</label>
              <div className="relative">
                <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
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
                    className="text-slate-400"
                  >
                    <circle cx="12" cy="12" r="10"></circle>
                    <line x1="12" y1="8" x2="12" y2="16"></line>
                    <line x1="8" y1="12" x2="16" y2="12"></line>
                  </svg>
                </div>
                <input
                  value={index}
                  onChange={(event) => setIndex(event.target.value)}
                  className="w-full pl-10 p-3 border border-slate-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 outline-none transition"
                  type="number"
                  placeholder="Enter a number..."
                />
              </div>
              <p className="mt-2 text-xs text-slate-500">Enter a positive integer to calculate its Fibonacci value</p>
            </div>

            <button
              className={`w-full px-5 py-3 bg-gradient-to-r from-blue-500 to-indigo-600 text-white rounded-lg font-medium transition-all ${
                loading ? "opacity-70 cursor-not-allowed" : "hover:from-blue-600 hover:to-indigo-700"
              }`}
              disabled={loading}
            >
              {loading ? (
                <span className="flex items-center justify-center">
                  <svg className="animate-spin -ml-1 mr-2 h-4 w-4 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                    <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4"></circle>
                    <path
                      className="opacity-75"
                      fill="currentColor"
                      d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
                    ></path>
                  </svg>
                  Processing...
                </span>
              ) : (
                "Calculate"
              )}
            </button>
          </form>
        </div>

        <div className="mt-8 bg-white rounded-xl border border-slate-200 shadow-sm overflow-hidden">
          <div className="p-6 border-b border-slate-200">
            <h2 className="text-xl font-semibold text-slate-800">Calculated Indices</h2>
            <p className="text-slate-500 mt-1">Previously calculated indices</p>
          </div>
          <div className="p-6">{renderSeenIndexes()}</div>
        </div>
      </div>

      <div>
        <div className="bg-white rounded-xl border border-slate-200 shadow-sm overflow-hidden">
          <div className="p-6 border-b border-slate-200">
            <h2 className="text-xl font-semibold text-slate-800">Results</h2>
            <p className="text-slate-500 mt-1">Fibonacci calculation results</p>
          </div>
          <div className="p-6">{renderValues()}</div>
        </div>

        <div className="mt-8 bg-white rounded-xl border border-slate-200 shadow-sm p-6 text-sm text-slate-700">
          <h3 className="text-lg font-semibold mb-2 text-slate-800">About Fibonacci</h3>
          <p className="mb-3">The Fibonacci sequence is a series of numbers where each number is the sum of the two preceding ones, starting from 0 and 1.</p>
          <div className="p-3 bg-slate-50 rounded-lg font-mono text-center">0, 1, 1, 2, 3, 5, 8, 13, 21, 34, ...</div>
        </div>
      </div>
    </div>
  );
};

export default Fib;
