class CreateChats < ActiveRecord::Migration[7.1]
  def change
    create_table :chats do |t|
      t.integer :chat_number , null: false , index: true
      t.string :app_token , null: false 
      t.integer :message_count , null: false

      t.timestamps
    end
  end
end
