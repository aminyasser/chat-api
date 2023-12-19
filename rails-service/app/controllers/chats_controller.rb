class ChatsController < ApplicationController
 # GET /apps/:app_token/chats
  def get_chats

      begin        
          @app = App.find_by!(app_token: params[:app_token])
      rescue ActiveRecord::RecordNotFound => e
          render json: {status: 'error' , message:  "app does't exist"} , status: 404
          return             
      end

      @chats = Chat.where(app_token: params[:app_token])

      if @chats.empty?
          render json: { status: 'success' , message: "there is no created chats yet." }, status: 404
      else
          render json: { status: 'success' , message: "chats fetched successfully" , chats: @chats.to_json(except: [:id, :app_token]) } , status: 200
      end
  end
 
end
