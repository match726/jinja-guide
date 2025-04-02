import { SlPlus, SlMinus } from 'react-icons/sl';

import { Button } from '@/components/ui/button';
import { FieldProps, FormProps } from '@/features/register/types/types';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';

interface RegisterFormProps {
  title?: string;
  sections: FormProps[];
  onValueChange: (sectionName: string, fieldId: number, target: string, value: string) => void;
  onAddField: (sectionName: string) => void;
  onRemoveField: (sectionName: string, fieldId: number) => void;
  onSubmit: (e: React.FormEvent) => void;
}

const RegisterForm = ({ title, sections, onValueChange, onAddField, onRemoveField, onSubmit }: RegisterFormProps) => {

  return (
    <div className="w-full max-w-lg bg-white rounded-lg shadow-xl overflow-hidden">
      <div className="bg-red-900 p-4 flex items-center justify-center">
        <h2 className="text-2xl font-bold text-white ml-2 font-serif">
          {title}
        </h2>
      </div>
      <form onSubmit={onSubmit} className="p-6 space-y-6">
        {sections.map((section) => (
          <div key={section.name}>
            <Label htmlFor={section.name} className="text-lg font-medium text-gray-700 font-serif">
              {section.title}
            </Label>
            {section.fields.map((field) => (
              <div key={field.id} className="flex items-end gap-2">
                <div className="flex-1">
                  {renderFormField(section, field, onValueChange, onAddField, onRemoveField)}
                </div>
              </div>
            ))}
          </div>
        ))}
        <Button className="w-full bg-red-900 hover:bg-red-800 text-white font-bold py-2 px-4 rounded-md transition duration-300 ease-in-out transform hover:scale-105 font-serif">
          登録
        </Button>
      </form>
    </div>
  )

};

const renderFormField = (
  section: FormProps,
  field: FieldProps,
  onValueChange: (sectionName: string, fieldId: number, target: string, value: string) => void,
  onAddField: (sectionName: string) => void,
  onRemoveField: (sectionName: string, fieldId: number) => void
) => {
  switch (section.type) {
    case "text":
      return (
        <div className="flex justify-center items-center gap-2">
          <Input
            id={`${section.name}-${field.id}`}
            value={field.content1}
            onChange={(e) => onValueChange(section.name, field.id, "content1", e.target.value)}
            placeholder={section.placeHolder1}
            className="w-full border-2 border-red-800 rounded-md p-2 font-serif"
          />
          {section.isMultiple && (
            <div>
              <SlPlus onClick={() => onAddField(section.name)} className="h-5 w-5" />
              {section.fields.length > 1 && (
                <SlMinus onClick={() => onRemoveField(section.name, field.id)} className="h-5 w-5" />
              )}
            </div>
          )}
        </div>
      );
    case "text2":
      return (
        <div className="flex justify-center items-center gap-2">
          <Input
            id={`${section.name}-${field.id}1`}
            value={field.content1}
            onChange={(e) => onValueChange(section.name, field.id, "content1", e.target.value)}
            placeholder={section.placeHolder1}
            className="w-full border-2 border-red-800 rounded-md p-2 font-serif"
          />
          <Input
            id={`${section.name}-${field.id}2`}
            value={field.content2}
            onChange={(e) => onValueChange(section.name, field.id, "content2", e.target.value)}
            placeholder={section.placeHolder2}
            className="w-full border-2 border-red-800 rounded-md p-2 font-serif"
          />
          {section.isMultiple == true
            ? <div>
                <SlPlus onClick={() => onAddField(section.name)} className="h-5 w-5" />
                {section.fields.length > 1 && (
                  <SlMinus onClick={() => onRemoveField(section.name, field.id)} className="h-5 w-5" />
                )}
              </div>
            : null
          }
        </div>
      )  
    case "select + text2":
      return (
        <div className="flex justify-center items-center gap-2">
          <Select value={field.seq} onValueChange={(value) => onValueChange(section.name, field.id, "seq", value)}>
            <SelectTrigger className="w-full w-max-md border-2 border-red-800 rounded-md p-2 font-serif">
            <SelectValue placeholder={`${section.title}を選択`} />
          </SelectTrigger>
          <SelectContent>
              {section.options?.map((option) => (
                <SelectItem key={option} value={option.slice(0,option.indexOf(":"))}>
                {option}
              </SelectItem>
            ))}
          </SelectContent>
        </Select>
          <Input
            id={`${section.name}-${field.id}1`}
            value={field.content1}
            onChange={(e) => onValueChange(section.name, field.id, "content1", e.target.value)}
            placeholder={section.placeHolder1}
            className="w-full w-max-md border-2 border-red-800 rounded-md p-2 font-serif"
          />
          <Input
            id={`${section.name}-${field.id}2`}
            value={field.content2}
            onChange={(e) => onValueChange(section.name, field.id, "content2", e.target.value)}
            placeholder={section.placeHolder2}
            className="w-full w-max-md border-2 border-red-800 rounded-md p-2 font-serif"
          />
          {section.isMultiple == true
            ? <div>
                <SlPlus onClick={() => onAddField(section.name)} className="h-5 w-5" />
                {section.fields.length > 1 && (
                  <SlMinus onClick={() => onRemoveField(section.name, field.id)} className="h-5 w-5" />
                )}
              </div>
            : null
          }
        </div>
      )
    default:
      return null;
  }
};

export { RegisterForm };