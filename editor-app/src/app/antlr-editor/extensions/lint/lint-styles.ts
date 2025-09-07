// Apply lint tooltip styles dynamically
export const applyLintStyles = () => {
  // Check if styles are already applied
  if (document.getElementById('lint-custom-styles')) {
    return;
  }

  const styleSheet = document.createElement('style');
  styleSheet.id = 'lint-custom-styles';
  styleSheet.textContent = `
    /* Lint tooltip styles - matching hover and autocomplete design */
    .cm-tooltip.cm-lint-tooltip {
      background: #2d3748 !important;
      border: 1px solid rgba(99, 179, 237, 0.2) !important;
      border-radius: 8px !important;
      padding: 12px !important;
      max-width: 450px !important;
      box-shadow: 0 4px 12px rgba(0, 0, 0, 0.3) !important;
      color: #e2e8f0 !important;
      font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Roboto', sans-serif !important;
      font-size: 13px !important;
      line-height: 1.5 !important;
    }

    .cm-diagnostic {
      padding: 8px 12px !important;
      border-radius: 4px !important;
      margin: 4px 0 !important;
      background: rgba(0, 0, 0, 0.2) !important;
    }

    .cm-diagnostic-error {
      border-left: 3px solid #ef4444 !important;
      background: rgba(239, 68, 68, 0.1) !important;
    }

    .cm-diagnostic-warning {
      border-left: 3px solid #f59e0b !important;
      background: rgba(245, 158, 11, 0.1) !important;
    }

    .cm-diagnostic-info {
      border-left: 3px solid #3b82f6 !important;
      background: rgba(59, 130, 246, 0.1) !important;
    }

    .cm-diagnosticText {
      color: #f7fafc !important;
      font-size: 13px !important;
      line-height: 1.6 !important;
    }

    .cm-diagnosticSource {
      color: #68d391 !important;
      font-size: 11px !important;
      opacity: 0.8 !important;
      margin-top: 4px !important;
      font-family: 'SF Mono', Monaco, 'Cascadia Code', monospace !important;
    }
  `;
  document.head.appendChild(styleSheet);
};
