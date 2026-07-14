import { useRef, useState } from "react";

interface ResumeUploaderProps {
  resume: File | null;
  onResumeChange: (file: File | null) => void;
}

export default function ResumeUploader({ resume, onResumeChange }: ResumeUploaderProps) {
  const fileInputRef = useRef<HTMLInputElement>(null);
  const [dragging, setDragging] = useState(false);

  const handleFile = (file: File | undefined) => {
    if (!file) return;
    if (file.type === "application/pdf") {
      onResumeChange(file);
    } else {
      alert("Please upload a PDF file.");
    }
  };

  return (
    <div>
      {!resume ? (
        <div
          className={`upload-area${dragging ? " dragging" : ""}`}
          onClick={() => fileInputRef.current?.click()}
          onDrop={(e) => { e.preventDefault(); setDragging(false); handleFile(e.dataTransfer.files?.[0]); }}
          onDragOver={(e) => { e.preventDefault(); setDragging(true); }}
          onDragLeave={() => setDragging(false)}
          role="button"
          tabIndex={0}
          onKeyDown={(e) => e.key === "Enter" && fileInputRef.current?.click()}
          aria-label="Upload resume PDF"
        >
          <h3>Drop the candidate's resume here</h3>
          <p>or click to browse — PDF only</p>
          <input
            ref={fileInputRef}
            type="file"
            accept=".pdf"
            onChange={(e) => handleFile(e.target.files?.[0])}
            style={{ display: "none" }}
          />
        </div>
      ) : (
        <div className="file-selected">
          <div className="file-badge">📄</div>
          <div className="file-info">
            <strong>{resume.name}</strong>
            <span>{(resume.size / 1024).toFixed(1)} KB · PDF</span>
          </div>
          <button
            className="btn-remove"
            onClick={() => onResumeChange(null)}
            aria-label="Remove file"
          >
            ✕
          </button>
          <input
            ref={fileInputRef}
            type="file"
            accept=".pdf"
            onChange={(e) => handleFile(e.target.files?.[0])}
            style={{ display: "none" }}
          />
        </div>
      )}
    </div>
  );
}
