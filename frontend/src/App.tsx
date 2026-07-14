import { useState } from "react";
import "./App.css";
import Header from "../components/Header";
import ResumeUploader from "../components/ResumeUploader";
import JobDescriptionInput from "../components/JobDescriptionInput";
import ResultsDisplay from "../components/ResultsDisplay";

interface AnalysisResult {
  strong_points: string[];
  weak_points: string[];
  conclusion: string;
  match_percentage: string;
}

export default function App() {
  const [resume, setResume] = useState<File | null>(null);
  const [jobDescription, setJobDescription] = useState("");
  const [loading, setLoading] = useState(false);
  const [result, setResult] = useState<AnalysisResult | null>(null);
  const [error, setError] = useState<string | null>(null);

  const canAnalyze = resume !== null && jobDescription.trim().length > 0 && !loading;

  const handleAnalyze = async () => {
    if (!resume || !jobDescription.trim()) return;
    setLoading(true);
    setError(null);
    setResult(null);

    try {
      const formData = new FormData();
      formData.append("resume", resume);
      formData.append("job", jobDescription.trim());

      const response = await fetch("http://localhost:8080/analyze", {
        method: "POST",
        body: formData,
      });

      const data = await response.json();

      if (!response.ok) {
        setError(data.error || "Something went wrong. Please try again.");
        return;
      }

      setResult(data as AnalysisResult);
    } catch {
      setError("Could not reach the server. Make sure the backend is running on port 8080.");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="app">
      <Header />

      <main className="main">
        <div className="card">
          <p className="card-label">Candidate Resume</p>
          <ResumeUploader resume={resume} onResumeChange={setResume} />
        </div>

        <div className="card">
          <p className="card-label">Your Job Description</p>
          <JobDescriptionInput value={jobDescription} onChange={setJobDescription} />
        </div>

        <button
          className="btn-analyze"
          onClick={handleAnalyze}
          disabled={!canAnalyze}
          aria-busy={loading}
        >
          {loading ? (
            <><span className="spinner" /> Analyzing…</>
          ) : (
            "Screen Candidate"
          )}
        </button>

        {error && <div className="error-banner">{error}</div>}

        {result && (
          <div style={{ marginTop: 32 }}>
            <ResultsDisplay result={result} />
          </div>
        )}
      </main>
    </div>
  );
}
