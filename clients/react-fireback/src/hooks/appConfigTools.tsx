import React, { useState } from "react";

export interface AppConfig {
  remote?: string;
  theme?: string;
  interfaceLanguage?: string;
  audioInputDevice?: string;
  audioOutputDevice?: string;
  preferredHand?: string;
  videoInputDevice?: string;
  textEditorModule?: string;
}

export interface IAppConginContext {
  //   setConfig: (config: AppConfig) => void;
  patchConfig: (config: AppConfig) => void;
  config: AppConfig;
}

export const AppConfigContext = React.createContext<IAppConginContext>({
  //   setConfig() {},
  patchConfig() {},
  config: {},
});

function getUserConfig() {
  const userConfig = localStorage.getItem("app_config2");
  if (!userConfig) {
    return;
  }

  try {
    const cnf = JSON.parse(userConfig);
    if (cnf) {
      return { ...cnf };
    }
    return {};
  } catch (error) {}
  return {};
}

const appStorageConfig = getUserConfig();

export function AppConfigProvider({
  children,
  initialConfig,
}: {
  children: React.ReactNode;
  initialConfig: AppConfig;
}) {
  const [config, setConfig] = useState<AppConfig>({
    ...initialConfig,
    ...appStorageConfig,
  });
  const patchConfig = (config: Partial<AppConfig>) => {
    setConfig((c) => {
      const newConf = {
        ...c,
        ...config,
      };

      localStorage.setItem("app_config2", JSON.stringify(newConf));
      return newConf;
    });
  };

  return (
    <AppConfigContext.Provider value={{ config, patchConfig }}>
      {children}
    </AppConfigContext.Provider>
  );
}
