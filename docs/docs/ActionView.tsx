
const cliPrefix = 'fireback'
const urlPrefix = 'http://localhost:4500'

const ActionView = ({ cli, url, method} : {cli: string, url: string, method?: string}) => {
  // Combine CLI prefix if provided
  const fullCli = cliPrefix ? `${cliPrefix} ${cli}` : cli;
  // Combine URL prefix if provided
  const fullUrl = urlPrefix ? `${urlPrefix}${url}` : url;
  // Construct curl command
  const curlCommand = `curl -X ${(method || 'get').toUpperCase()} ${fullUrl}`;

  return (
    <div style={{ marginBottom: "1.5rem" }}>
      {cli ? <div>
        <strong>CLI:</strong>
        <pre>
          <code className="language-bash">{fullCli}</code>
        </pre>
      </div> : null}
      {url ? <div>
        <strong>cURL:</strong>
        <pre>
          <code className="language-bash">{curlCommand}</code>
        </pre>
      </div> : null}
    </div>
  );
};

export default ActionView;