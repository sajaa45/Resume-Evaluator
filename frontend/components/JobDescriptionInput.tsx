interface JobDescriptionInputProps {
  value: string;
  onChange: (value: string) => void;
}

export default function JobDescriptionInput({ value, onChange }: JobDescriptionInputProps) {
  return (
    <div>
      <textarea
        className="textarea"
        value={value}
        onChange={(e) => onChange(e.target.value)}
        placeholder="Paste your job description here — required skills, experience level, responsibilities..."
        aria-label="Your job description"
      />
      <p className="char-count">{value.length} characters</p>
    </div>
  );
}
