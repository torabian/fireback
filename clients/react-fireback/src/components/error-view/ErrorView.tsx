import React from "react";

export const ErrorsView = ({
  errors,
  error,
}: {
  error?: any;
  errors?: any;
}) => {
  if (!error && !errors) {
    return null;
  }

  let errorList: any = {};

  if (error && error.errors) {
    errorList = error.errors;
  } else if (errors) {
    errorList = errors;
  }

  const keys = Object.keys(errorList);

  // if (keys?.length === 0 && !(error?.title || error?.message)) {
  //   return null;
  // }

  return (
    <div style={{ minHeight: "30px" }}>
      {errors.form && (
        <div className="with-fade-in" style={{ color: "red" }}>
          {errors.form}
        </div>
      )}
      {errorList.length && (
        <div>
          {(error?.title || error?.message) && (
            <span>{error?.title || error?.message}</span>
          )}
          {keys.map((key: string) => {
            return (
              <div key={key}>
                <span>&bull; {(errorList as any)[key]}</span>
              </div>
            );
          })}
        </div>
      )}
    </div>
  );
};
