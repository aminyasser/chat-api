Rails.application.routes.draw do
  # Define your application routes per the DSL in https://guides.rubyonrails.org/routing.html

  # Reveal health status on /up that returns 200 if the app boots with no exceptions, otherwise 500.
  # Can be used by load balancers and uptime monitors to verify that the app is live.
  get "up" => "rails/health#show", as: :rails_health_check

  get "/apps", to: "apps#index"
  post "/apps", to: "apps#create"
  get "/apps/:app_token", to: "apps#show"

  get "/apps/:app_token/chats", to: "chats#get_chats"
  

end
