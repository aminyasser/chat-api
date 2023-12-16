class AppsController < ApplicationController
  # GET /apps
  def index
    @apps = App.all
    render json: @apps.as_json(except: [:id])
  end

  # GET /apps/:app_token
  def show
    @app = App.find_by!(app_token: params[:app_token])
    render json: @app.as_json(except: [:id])
  end

  # POST /apps 
  def create
    
    @app = App.create(app_params.merge(:chat_count => 0))
    if @app.valid?
      render json: { status:'success' , message:'App created successfully' , data: {app_token: @app.app_token }}, status: :created  
    else
      render json: {status: 'error' , message: @app.errors.objects.first.full_message}, status:  :unprocessable_entity  
    end
  end

  def app_params
    params.permit(:name)
  end

end
