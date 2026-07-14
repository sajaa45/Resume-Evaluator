import { useEffect, useState } from "react";

interface AnalysisResult {
  strong_points: string[];
  weak_points: string[];
  conclusion: string;
  match_percentage: string;
}

function parsePercent(raw: string): number {
  const n = parseInt(raw.replace(/[^0-9]/g, ""), 10);
  return isNaN(n) ? 0 : Math.min(100, Math.max(0, n));
}

function scoreColor(pct: number): string {
  if (pct >= 70) return "#2d8a4e";
  if (pct >= 45) return "#c47d00";
  return "#c0392b";
}

function scoreLabel(pct: number): string {
  if (pct >= 70) return "Strong candidate — meets most of your requirements.";
  if (pct >= 45) return "Partial match — some gaps against your requirements.";
  return "Weak match — significant gaps found.";
}

function ScoreRing({ pct }: { pct: number }) {
  const [animated, setAnimated] = useState(0);
  const r = 46;
  const circ = 2 * Math.PI * r;
  const color = scoreColor(pct);

  useEffect(() => {
    const t = setTimeout(() => setAnimated(pct), 80);
    return () => clearTimeout(t);
  }, [pct]);

  const offset = circ - (animated / 100) * circ;

  return (
    <div className="score-ring">
      <svg width="110" height="110" viewBox="0 0 110 110" aria-hidden="true">
        <circle className="score-ring-track" cx="55" cy="55" r={r} />
        <circle
          className="score-ring-bar"
          cx="55" cy="55" r={r}
          stroke={color}
          strokeDasharray={circ}
          strokeDashoffset={offset}
        />
      </svg>
      <div className="score-center">
        <span className="score-number" style={{ color }}>{pct}%</span>
        <span className="score-lbl">match</span>
      </div>
    </div>
  );
}

export default function ResultsDisplay({ result }: { result: AnalysisResult }) {
  const pct = parsePercent(result.match_percentage);

  return (
    <div className="results">
      <div>
        <h2 className="results-heading">Screening Results</h2>
        <p className="results-sub">Here's how this candidate's resume matches your job requirements.</p>
      </div>

      {/* Match score */}
      <div className="score-card">
        <ScoreRing pct={pct} />
        <div className="score-info">
          <h3>Match Score</h3>
          <p>{scoreLabel(pct)}</p>
        </div>
      </div>

      {/* Conclusion */}
      <div className="conclusion-card">
        <p className="conclusion-label">Conclusion</p>
        <p className="conclusion-value">{result.conclusion || "—"}</p>
      </div>

      {/* Points */}
      <div className="points-grid">
        <div className="points-card green">
          <div className="points-header">
            <span className="points-header-title">Strong Points</span>
            <span className="points-count">{result.strong_points?.length ?? 0}</span>
          </div>
          <ul className="points-list">
            {(result.strong_points ?? []).length > 0
              ? (result.strong_points ?? []).map((pt, i) => (
                  <li key={i} className="point-item">
                    <span className="point-dot" />
                    {pt}
                  </li>
                ))
              : <li className="points-empty">None identified.</li>
            }
          </ul>
        </div>

        <div className="points-card red">
          <div className="points-header">
            <span className="points-header-title">Weak Points</span>
            <span className="points-count">{result.weak_points?.length ?? 0}</span>
          </div>
          <ul className="points-list">
            {(result.weak_points ?? []).length > 0
              ? (result.weak_points ?? []).map((pt, i) => (
                  <li key={i} className="point-item">
                    <span className="point-dot" />
                    {pt}
                  </li>
                ))
              : <li className="points-empty">None identified.</li>
            }
          </ul>
        </div>
      </div>
    </div>
  );
}
