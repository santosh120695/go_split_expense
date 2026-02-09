import { X } from "lucide-react"
import React from "react"

interface CreateGroupModalProps {
  isOpen: boolean
  onClose: () => void
  formData: { name: string; description: string }
  onInputChange: (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => void
  onSubmit: (e: React.FormEvent) => void
}

function CreateGroupModal({
  isOpen,
  onClose,
  formData,
  onInputChange,
  onSubmit,
}: CreateGroupModalProps) {
  if (!isOpen) return null

  return (
    <>
      {isOpen && (
        <div 
          className="fixed inset-0 z-40 transition-opacity duration-200"
          onClick={onClose}
        />
      )}
      {isOpen && (
        <div className="fixed inset-0 z-50 flex items-center justify-center p-4 pointer-events-none bg-(--card)">
          <div 
            className="bg-card rounded-lg border border-[#C5C3C3] p-8 w-full max-w-md shadow-lg pointer-events-auto"
            onClick={(e) => e.stopPropagation()}
          >
        {/* Modal Header */}
        <div className="flex items-center justify-between mb-6">
          <h2 className="text-2xl font-bold text-foreground">Create Group</h2>
          <button
            onClick={onClose}
            className="p-1 hover:bg-accent rounded-lg transition-colors duration-200"
          >
            <X size={24} className="text-muted-foreground" />
          </button>
        </div>

        {/* Form */}
        <form onSubmit={onSubmit} className="space-y-4">
          <div>
            <label htmlFor="name" className="block text-sm font-medium text-foreground mb-2">
              Group Name *
            </label>
            <input
              type="text"
              id="name"
              name="name"
              value={formData.name}
              onChange={onInputChange}
              placeholder="e.g., House Bills"
              className="w-full px-4 py-2 bg-background border border-[#C5C3C3] rounded-lg text-foreground placeholder-muted-foreground focus:outline-none focus:ring-2 focus:ring-primary transition-all duration-200"
              required
            />
          </div>

          <div>
            <label htmlFor="description" className="block text-sm font-medium text-foreground mb-2">
              Description
            </label>
            <textarea
              id="description"
              name="description"
              value={formData.description}
              onChange={onInputChange}
              placeholder="What is this group for?"
              rows={4}
              className="w-full px-4 py-2 bg-background border border-[#C5C3C3] rounded-lg text-foreground placeholder-muted-foreground focus:outline-none focus:ring-2 focus:ring-primary transition-all duration-200 resize-none"
            />
          </div>

          {/* Form Actions */}
          <div className="flex gap-3 pt-4">
            <button
              type="button"
              onClick={onClose}
              className="flex-1 px-4 py-2  text-(--danger) border border-(--danger) rounded-lg hover:bg-opacity-80 transition-colors duration-200 font-medium"
            >
              Cancel
            </button>
            <button
              type="submit"
              className="flex-1 px-4 py-2 border-(--primary) border text-(--primary) rounded-lg hover:bg-opacity-90 transition-all duration-200 font-medium"
            >
              Create
            </button>
          </div>
        </form>
          </div>
        </div>
      )}
    </>
  )
}

export default CreateGroupModal
