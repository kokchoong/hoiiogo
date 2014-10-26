package hoiiogo

import (
  "net/url"
)

func IvrDial(app *App, msg, dest, callerId, tag, notifyURL string) (*Txn, error) { 
  params := url.Values{}
  params.Set("app_id", app.AppId())
  params.Set("access_token", app.AccessToken())
  params.Set("msg", msg)
  params.Set("dest", dest)
  params.Set("caller_id", callerId)
  params.Set("tag", tag)
  params.Set("notify_url", notifyURL)
  return app.Do(params, IVR_DIAL)
}

func IvrPlay(app *App, session, msg, tag, notifyURL string) (*Txn, error) { 
  params := url.Values{}
  params.Set("app_id", app.AppId())
  params.Set("access_token", app.AccessToken())
  params.Set("session", session)
  params.Set("msg", msg)
  params.Set("tag", tag)
  params.Set("notify_url", notifyURL)
  return app.Do(params, IVR_PLAY)
}

func IvrGather(app *App, session, msg, maxDigits, timeout, attempts, tag, notifyURL string) (*Txn, error) { 
  params := url.Values{}
  params.Set("app_id", app.AppId())
  params.Set("access_token", app.AccessToken())
  params.Set("session", session)
  params.Set("msg", msg)
  params.Set("max_digits", maxDigits)
  params.Set("timeout", timeout)
  params.Set("attempts", attempts)
  params.Set("tag", tag)
  params.Set("notify_url", notifyURL)
  return app.Do(params, IVR_GATHER)
}

func IvrRecord(app *App, session, msg, maxDuration, tag, notifyURL string) (*Txn, error) { 
  params := url.Values{}
  params.Set("app_id", app.AppId())
  params.Set("access_token", app.AccessToken())
  params.Set("session", session)
  params.Set("msg", msg)
  params.Set("max_duration", maxDuration)
  params.Set("tag", tag)
  params.Set("notify_url", notifyURL)
  return app.Do(params, IVR_RECORD)
}

func IvrMonitor(app *App, session, msg, tag, notifyURL, monitorNotifyURl string) (*Txn, error) { 
  params := url.Values{}
  params.Set("app_id", app.AppId())
  params.Set("access_token", app.AccessToken())
  params.Set("session", session)
  params.Set("msg", msg)
  params.Set("tag", tag)
  params.Set("notify_url", notifyURL)
  params.Set("monitor_notify_url", monitorNotifyURl)
  return app.Do(params, IVR_MONITOR)
}

func IvrHold(app *App, session, msg string) (*Txn, error) { 
  params := url.Values{}
  params.Set("app_id", app.AppId())
  params.Set("access_token", app.AccessToken())
  params.Set("session", session)
  params.Set("msg", msg)
  return app.Do(params, IVR_HOLD)
}

func IvrUnHold(app *App, session, msg, tag, notifyURL string) (*Txn, error) { 
  params := url.Values{}
  params.Set("app_id", app.AppId())
  params.Set("access_token", app.AccessToken())
  params.Set("session", session)
  params.Set("msg", msg)
  params.Set("tag", tag)
  params.Set("notify_url", notifyURL)
  return app.Do(params, IVR_UNHOLD)
}

func IvrTransfer(app *App, session, msg, dest, callerId, onFailure, tag, notifyURL string) (*Txn, error) { 
  params := url.Values{}
  params.Set("app_id", app.AppId())
  params.Set("access_token", app.AccessToken())
  params.Set("session", session)
  params.Set("msg", msg)
  params.Set("dest", dest)
  params.Set("caller_id", callerId)
  params.Set("on_failure", onFailure)
  params.Set("tag", tag)
  params.Set("notify_url", notifyURL)
  return app.Do(params, IVR_TRANSFER)
}

func IvrHangUp(app *App, session, msg, tag, notifyURL string) (*Txn, error) { 
  params := url.Values{}
  params.Set("app_id", app.AppId())
  params.Set("access_token", app.AccessToken())
  params.Set("session", session)
  params.Set("msg", msg)
  params.Set("tag", tag)
  params.Set("notify_url", notifyURL)
  return app.Do(params, IVR_HANGUP)
}

func IvrForcedHangUp(app *App, session, tag, notifyURL string) (*Txn, error) { 
  params := url.Values{}
  params.Set("app_id", app.AppId())
  params.Set("access_token", app.AccessToken())
  params.Set("session", session)
  params.Set("force", "enabled")
  params.Set("tag", tag)
  params.Set("notify_url", notifyURL)
  return app.Do(params, IVR_HANGUP)
}

